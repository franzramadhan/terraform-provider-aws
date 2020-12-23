package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecrpublic"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAwsEcrPublicRegistryCatalogData() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsEcrPublicRegistryCatalogDataCreate,
		Read:   resourceAwsEcrPublicRegistryCatalogDataRead,
		Update: resourceAwsEcrPublicRegistryCatalogDataUpdate,
		Delete: resourceAwsEcrPublicRegistryCatalogDataDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAwsEcrPublicRegistryCatalogDataCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).ecrpublicconn

	input := &ecrpublic.PutRegistryCatalogDataInput{
		DisplayName: aws.String(d.Get("display_name").(string)),
	}

	_, err := conn.PutRegistryCatalogData(input)

	if err != nil {
		return fmt.Errorf("error changing ECR Public Registry catalog data: %s", err)
	}

	return resourceAwsEcrPublicRegistryCatalogDataRead(d, meta)
}

func resourceAwsEcrPublicRegistryCatalogDataRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).ecrpublicconn

	var out *ecrpublic.GetRegistryCatalogDataOutput
	input := &ecrpublic.GetRegistryCatalogDataInput{}

	out, err := conn.GetRegistryCatalogData(input)

	if err != nil {
		return fmt.Errorf("error reading ECR Public Registry catalog data: %s", err)
	}

	registryCatalogData := out.RegistryCatalogData

	d.Set("display_name", registryCatalogData.DisplayName)

	return nil
}

func resourceAwsEcrPublicRegistryCatalogDataUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).ecrpublicconn

	if d.HasChange("display_name") {
		input := &ecrpublic.PutRegistryCatalogDataInput{
			DisplayName: aws.String(d.Get("display_name").(string)),
		}

		_, err := conn.PutRegistryCatalogData(input)

		if err != nil {
			return fmt.Errorf("error changing ECR Public Registry catalog data: %s", err)
		}
	}

	return resourceAwsEcrPublicRegistryCatalogDataRead(d, meta)
}

func resourceAwsEcrPublicRegistryCatalogDataDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}