package week

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWeeks() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceWeeksRead,
		CreateContext: resourceWeeksCreate,
		DeleteContext: resourceWeeksDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"week_number": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"year": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func resourceWeeksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceWeeksCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	year, week_number := time.Now().ISOWeek() // year and week_number
	err := d.Set("week_number", week_number)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("year", year)
	if err != nil {
		return diag.FromErr(err)
	}

	week_string := strconv.Itoa(week_number)
	year_string := strconv.Itoa(year)
	id := fmt.Sprintf("%s%s", year_string, week_string)
	d.SetId(id)
	return nil
}

func resourceWeeksDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
