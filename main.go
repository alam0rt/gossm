package main

import (
	"flag"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	regionPtr := flag.String("region", "ap-southeast-2", "The region in which to operate")
	pathPtr := flag.String("path", "/", "The path of the parameters in which to recurse")
	recursePtr := flag.Bool("recurse", false, "recurse through SSM hierarchy")
	decryptPtr := flag.Bool("decrypt", false, "decrypt the parameters")
	maxResultsPtr := flag.Int64("results", 10, "how many results to get")
	flag.Parse()
	//	fmt.Println("region: ", *regionPtr)
	//	fmt.Println("path: ", *pathPtr)
	//	fmt.Println("recurse: ", *recursePtr)
	//	fmt.Println("decrypt: ", *decryptPtr)
	//	fmt.Printf("\n\n")


	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ssm.New(sess, &aws.Config{
		Region: aws.String(*regionPtr),
	})

	getParameterPath(svc, pathPtr, recursePtr, decryptPtr, maxResultsPtr)
}

func exitProgram() {
	fmt.Println("ruh roh!")
	os.Exit(1)
}

func prompt(keys []string) string {
	prompt := promptui.Select{
		Label: "Choose parameter",
		Items: keys,
		
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Print(err)
	}
	return result
}

func getParameter(svc *ssm.SSM, name *string, decrypt *bool) {
	p, err := svc.GetParameter(&ssm.GetParameterInput{
		Name: name,
	})
	if err != nil {
		fmt.Println(err)
		exitProgram()
	}
	fmt.Print(*p.Parameter.Value)
}

func getParameterPath(svc *ssm.SSM, path *string, recurse *bool, decrypt *bool, results *int64) {
	p, err := svc.GetParametersByPath(&ssm.GetParametersByPathInput{
		Path:           path,
		Recursive:      recurse,
		WithDecryption: decrypt,
		MaxResults: results,
	})
	if err != nil {
		fmt.Print(err)
	}
	param := make([]string, *results)
	keys := param
	for i, pa := range p.Parameters {
		if p == nil {
			continue
		}
		param[i] = *pa.Name // add param to array
		keys = param[:i+1] // set slice as len of array
	}
	r := prompt(keys)
	getParameter(svc, &r, decrypt)

}
