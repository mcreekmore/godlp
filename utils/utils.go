package utils

import (
	"fmt"
	"godlp/embed"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly/v2"
)

func MoveFiles(source, dest string) error {
	files, err := filepath.Glob(filepath.Join(source, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println("file", file)
		_, fileName := filepath.Split(file)
		inputFile, err := os.Open(file)
		if err != nil {
			fmt.Printf("Couldn't open source file: %v\n", err)
		}
		outputFile, err := os.Create(dest + "/" + fileName)
		if err != nil {
			inputFile.Close()
			fmt.Printf("Couldn't open dest file: %v\n", err)
			return err
		}
		defer outputFile.Close()
		_, err = io.Copy(outputFile, inputFile)
		inputFile.Close()
		if err != nil {
			fmt.Printf("Writing to output file failed: %v\n", err)
			return err
		}
		// The copy was successful, so now delete the original file
		err = os.Remove(file)
		if err != nil {
			fmt.Printf("Failed removing original file: %v\n", err)
			return err
		}
	}

	return nil
}

func ExtractArtistNameFromFile(tempDir string) (string, error) {
	files, err := filepath.Glob(filepath.Join(tempDir, "*"))
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", fmt.Errorf("no files found in temp directory")
	}

	// Assuming the first file in tempDir contains the downloaded file name
	fileName := filepath.Base(files[0])

	// Split the filename using the "-" delimiter
	parts := strings.Split(fileName, "-")
	if len(parts) < 2 {
		return "", fmt.Errorf("unable to extract artist name from the filename")
	}

	// Extract the artist name and trim whitespaces
	artistName := strings.TrimSpace(parts[0])

	return artistName, nil
}

func ChangeAlbumNameWithFFmpeg(directory, albumName string) error {
	files, err := filepath.Glob(filepath.Join(directory, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		ffmpegArgs := []string{
			"-y",
			"-i", file,
			"-metadata", fmt.Sprintf("album=%s", albumName),
			"-c", "copy",
			file + "_temp.mp3",
		}
		err := embed.ExecuteFfmpeg(ffmpegArgs)
		if err != nil {
			return err
		}

		// Replace the original file with the one containing updated metadata
		err = os.Rename(file+"_temp.mp3", file)
		if err != nil {
			return err
		}
	}

	return nil
}

func ScrapeSoundcloudAlbumTitle(url string) (string, error) {
	var title string

	c := colly.NewCollector()

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	// c.OnHTML("*", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })
	// c.OnHTML("article header h1[itemprop=name]", func(e *colly.HTMLElement) {
	// 	// Find the <a> within the <h1> element
	// 	a := e.ChildAttr("a", "href")
	// 	title = strings.TrimPrefix(a, "/")

	// 	// Print the full title for debugging
	// 	fmt.Println("Full Title:", title)
	// })
	// c.OnHTML("article header h1.soundTitle__title", func(e *colly.HTMLElement) {
	// 	// Extract text content of the h1 element
	// 	title = e.Text

	// 	// Print the full title for debugging
	// 	fmt.Println("Full Title:", title)
	// })
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})
	c.OnHTML("h1.sounTitle__title", func(e *colly.HTMLElement) {
		// printing all URLs associated with the a links in the page
		fmt.Println(e)
		// fmt.Println(e.ChildText("a"))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(url)
	if err != nil {
		return "", err
	}

	c.Wait()

	fmt.Println("title: ", title)
	return title, nil
}
