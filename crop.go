/* Simple command line program to crop images. 
Usage: 
$ go run crop.go --height=5000 --width=7000 <PATH TO>/CAT1.JPG <PATH TO>/CAT2.PNG
The cropped images will be placed in the same directory as the original images 
with the file names being
cropped_<original_file_name>.<original_extension>
*/

package main
import (
    "bufio"
    "bytes"
    "flag"
    "fmt"
    "github.com/oliamb/cutter"
    "image"
    "image/jpeg"
    "image/png"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
)

func validateImageType(imageType string) bool {
    recognizedImagesTypes := []string{"image/jpeg","image/png"}
    for _,t:= range recognizedImageTypes{
        if imageType == t{
            return true
        }
    }
    return false
}
                
