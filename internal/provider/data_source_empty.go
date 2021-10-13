package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEmpty() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Empty.",

		ReadContext: dataSourceEmptyRead,

		Schema: map[string]*schema.Schema{
			"empty_attribute": {
				// This description is used by the documentation generator and the language server.
				Description: "Empty.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceEmptyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)
	return nil
	//return diag.Errorf("not implemented")
}
