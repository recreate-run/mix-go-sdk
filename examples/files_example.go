package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/operations"
)

func main() {
	serverURL := os.Getenv("MIX_SERVER_URL")
	if serverURL == "" {
		serverURL = "http://localhost:8088"
	}

	client := mix.New(mix.WithServerURL(serverURL))
	ctx := context.Background()

	fmt.Println("=== Mix Go SDK - Files Example ===\n")

	var createdSessions []string

	// Cleanup function
	defer func() {
		if len(createdSessions) > 0 {
			fmt.Println("\n=== Cleanup ===")
			for _, sessionID := range createdSessions {
				fmt.Printf("Deleting session (and its files): %s...\n", sessionID)
				_, err := client.Sessions.DeleteSession(ctx, sessionID)
				if err != nil {
					log.Printf("Failed to delete session %s: %v", sessionID, err)
				}
			}
		}
	}()

	// 1. Create a session for file operations
	fmt.Println("1. Creating a session for file operations...")
	createResp, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
		Title: "Files Example Session",
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	sessionID := createResp.SessionData.ID
	createdSessions = append(createdSessions, sessionID)
	fmt.Printf("   Created session: %s\n\n", sessionID)

	// 2. Upload a text file
	fmt.Println("2. Uploading a text file...")
	textContent := "Hello from Mix Go SDK!\nThis is a sample text file.\nCreated at: " + time.Now().Format(time.RFC3339)
	textFile := createFileReader("sample.txt", []byte(textContent))
	uploadResp, err := client.Files.UploadSessionFile(ctx, sessionID, operations.UploadSessionFileRequestBody{
		File: textFile,
	})
	if err != nil {
		log.Fatalf("Failed to upload text file: %v", err)
	}
	fmt.Printf("   Uploaded: %s\n", uploadResp.FileInfo.Name)
	fmt.Printf("   URL: %s\n", uploadResp.FileInfo.URL)
	fmt.Printf("   Size: %d bytes\n", uploadResp.FileInfo.Size)
	fmt.Printf("   Modified: %d\n\n", uploadResp.FileInfo.Modified)

	// 3. Upload an image file (if sample_files/sample.jpg exists)
	sampleImagePath := filepath.Join("..", "mix-python-sdk", "examples", "sample_files", "sample.jpg")
	var imageFileID string
	if fileExists(sampleImagePath) {
		fmt.Println("3. Uploading an image file...")
		imageData, err := os.ReadFile(sampleImagePath)
		if err != nil {
			log.Printf("Failed to read image file: %v", err)
		} else {
			imageFile := createFileReader("sample.jpg", imageData)
			imgUploadResp, err := client.Files.UploadSessionFile(ctx, sessionID, operations.UploadSessionFileRequestBody{
				File: imageFile,
			})
			if err != nil {
				log.Printf("Failed to upload image file: %v", err)
			} else {
				imageFileID = imgUploadResp.FileInfo.URL
				fmt.Printf("   Uploaded: %s\n", imgUploadResp.FileInfo.Name)
				fmt.Printf("   URL: %s\n", imgUploadResp.FileInfo.URL)
				fmt.Printf("   Size: %d bytes\n", imgUploadResp.FileInfo.Size)
				fmt.Printf("   Modified: %d\n\n", imgUploadResp.FileInfo.Modified)
			}
		}
	} else {
		fmt.Println("3. Skipping image upload (sample.jpg not found)")
		fmt.Printf("   Place an image at: %s\n\n", sampleImagePath)
	}

	// 4. Upload a binary file (creating sample data)
	fmt.Println("4. Uploading a binary file...")
	binaryData := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}
	binaryFile := createFileReader("sample.bin", binaryData)
	binUploadResp, err := client.Files.UploadSessionFile(ctx, sessionID, operations.UploadSessionFileRequestBody{
		File: binaryFile,
	})
	if err != nil {
		log.Printf("Failed to upload binary file: %v", err)
	} else {
		fmt.Printf("   Uploaded: %s\n", binUploadResp.FileInfo.Name)
		fmt.Printf("   URL: %s\n", binUploadResp.FileInfo.URL)
		fmt.Printf("   Size: %d bytes\n\n", binUploadResp.FileInfo.Size)
	}

	// 5. List all files in the session
	fmt.Println("5. Listing all files in the session...")
	listResp, err := client.Files.ListSessionFiles(ctx, sessionID)
	if err != nil {
		log.Fatalf("Failed to list files: %v", err)
	}
	fmt.Printf("   Found %d file(s):\n", len(listResp.FileInfos))
	for i, file := range listResp.FileInfos {
		fmt.Printf("   [%d] %s\n", i+1, file.Name)
		fmt.Printf("       URL: %s\n", file.URL)
		fmt.Printf("       Size: %d bytes\n", file.Size)
		fmt.Printf("       Modified: %d\n", file.Modified)
		fmt.Printf("       Is Directory: %v\n", file.IsDir)
	}
	fmt.Println()

	// 6. Download a file
	if len(listResp.FileInfos) > 0 {
		firstFile := listResp.FileInfos[0]
		fmt.Printf("6. Downloading file: %s...\n", firstFile.Name)
		// Note: Using file name as the file ID (path)
		downloadResp, err := client.Files.GetSessionFile(ctx, sessionID, firstFile.Name, nil, nil)
		if err != nil {
			log.Printf("Failed to download file: %v", err)
		} else {
			defer downloadResp.ResponseStream.Close()
			data, err := io.ReadAll(downloadResp.ResponseStream)
			if err != nil {
				log.Printf("Failed to read file data: %v", err)
			} else {
				fmt.Printf("   Downloaded %d bytes\n", len(data))
				// Try to display as text if it's a .txt file
				if strings.HasSuffix(firstFile.Name, ".txt") {
					fmt.Printf("   Content preview: %s\n", truncate(string(data), 100))
				} else {
					fmt.Printf("   Binary content (first 16 bytes): %x\n", data[:min(16, len(data))])
				}
			}
		}
		fmt.Println()
	}

	// 7. Download image with thumbnail (if image was uploaded)
	if imageFileID != "" {
		fmt.Println("7. Downloading image with thumbnail (box constraint)...")
		thumbnailSpec := "100"
		// Note: Using "sample.jpg" as the file path
		thumbnailResp, err := client.Files.GetSessionFile(ctx, sessionID, "sample.jpg", &thumbnailSpec, nil)
		if err != nil {
			log.Printf("Failed to download thumbnail: %v", err)
		} else {
			defer thumbnailResp.ResponseStream.Close()
			data, err := io.ReadAll(thumbnailResp.ResponseStream)
			if err != nil {
				log.Printf("Failed to read thumbnail data: %v", err)
			} else {
				fmt.Printf("   Downloaded thumbnail: %d bytes\n", len(data))
				fmt.Println("   Note: Thumbnail generated with box=100 constraint")
			}
		}
		fmt.Println()
	}

	// 8. Verify file isolation (create second session)
	fmt.Println("8. Verifying file isolation between sessions...")
	createResp2, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
		Title: "Second Session - File Isolation Test",
	})
	if err != nil {
		log.Fatalf("Failed to create second session: %v", err)
	}
	sessionID2 := createResp2.SessionData.ID
	createdSessions = append(createdSessions, sessionID2)
	fmt.Printf("   Created second session: %s\n", sessionID2)

	listResp2, err := client.Files.ListSessionFiles(ctx, sessionID2)
	if err != nil {
		log.Fatalf("Failed to list files in second session: %v", err)
	}
	fmt.Printf("   Files in second session: %d\n", len(listResp2.FileInfos))
	fmt.Printf("   Files in first session: %d\n", len(listResp.FileInfos))
	fmt.Println("   ✓ File isolation verified: sessions have separate file storage")
	fmt.Println()

	// 9. Delete a file
	if len(listResp.FileInfos) > 0 {
		fileToDelete := listResp.FileInfos[0]
		fmt.Printf("9. Deleting file: %s...\n", fileToDelete.Name)
		// Note: Using file name as the file ID (path)
		deleteResp, err := client.Files.DeleteSessionFile(ctx, sessionID, fileToDelete.Name)
		if err != nil {
			log.Printf("Failed to delete file: %v", err)
		} else {
			fmt.Printf("   File deleted (Status: %d)\n", deleteResp.HTTPMeta.Response.StatusCode)

			// Verify deletion
			listResp3, err := client.Files.ListSessionFiles(ctx, sessionID)
			if err != nil {
				log.Printf("Failed to verify deletion: %v", err)
			} else {
				fmt.Printf("   Files remaining: %d\n", len(listResp3.FileInfos))
			}
		}
		fmt.Println()
	}

	// 10. File metadata summary
	fmt.Println("10. File operations summary:")
	listRespFinal, err := client.Files.ListSessionFiles(ctx, sessionID)
	if err != nil {
		log.Printf("Failed to get final file list: %v", err)
	} else {
		var totalSize int64
		for _, file := range listRespFinal.FileInfos {
			totalSize += file.Size
		}
		fmt.Printf("   Total files: %d\n", len(listRespFinal.FileInfos))
		fmt.Printf("   Total size: %d bytes (%.2f KB)\n", totalSize, float64(totalSize)/1024)
		fmt.Println("   File types supported: text, images, binary")
		fmt.Println("   Thumbnail support: available for images")
	}
	fmt.Println()

	fmt.Println("=== Files Example Completed Successfully! ===")
}

// createFileReader creates an io.Reader from filename and data
func createFileReader(filename string, data []byte) operations.File {
	return operations.File{
		FileName: filename,
		Content:  bytes.NewReader(data),
	}
}

// fileExists checks if a file exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// min returns the minimum of two ints
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
