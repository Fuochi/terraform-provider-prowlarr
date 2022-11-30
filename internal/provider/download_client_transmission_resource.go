package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/devopsarr/terraform-provider-sonarr/tools"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golift.io/starr/prowlarr"
)

const (
	downloadClientTransmissionResourceName   = "download_client_transmission"
	DownloadClientTransmissionImplementation = "Transmission"
	DownloadClientTransmissionConfigContrat  = "TransmissionSettings"
	DownloadClientTransmissionProtocol       = "torrent"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DownloadClientTransmissionResource{}
var _ resource.ResourceWithImportState = &DownloadClientTransmissionResource{}

func NewDownloadClientTransmissionResource() resource.Resource {
	return &DownloadClientTransmissionResource{}
}

// DownloadClientTransmissionResource defines the download client implementation.
type DownloadClientTransmissionResource struct {
	client *prowlarr.Prowlarr
}

// DownloadClientTransmission describes the download client data model.
type DownloadClientTransmission struct {
	Tags             types.Set    `tfsdk:"tags"`
	Name             types.String `tfsdk:"name"`
	Host             types.String `tfsdk:"host"`
	URLBase          types.String `tfsdk:"url_base"`
	Username         types.String `tfsdk:"username"`
	Password         types.String `tfsdk:"password"`
	TvCategory       types.String `tfsdk:"tv_category"`
	TvDirectory      types.String `tfsdk:"tv_directory"`
	RecentTvPriority types.Int64  `tfsdk:"recent_tv_priority"`
	OlderTvPriority  types.Int64  `tfsdk:"older_tv_priority"`
	Priority         types.Int64  `tfsdk:"priority"`
	Port             types.Int64  `tfsdk:"port"`
	ID               types.Int64  `tfsdk:"id"`
	AddPaused        types.Bool   `tfsdk:"add_paused"`
	UseSsl           types.Bool   `tfsdk:"use_ssl"`
	Enable           types.Bool   `tfsdk:"enable"`
}

func (d DownloadClientTransmission) toDownloadClient() *DownloadClient {
	return &DownloadClient{
		Tags:             d.Tags,
		Name:             d.Name,
		Host:             d.Host,
		URLBase:          d.URLBase,
		Username:         d.Username,
		Password:         d.Password,
		TvCategory:       d.TvCategory,
		TvDirectory:      d.TvDirectory,
		RecentTvPriority: d.RecentTvPriority,
		OlderTvPriority:  d.OlderTvPriority,
		Priority:         d.Priority,
		Port:             d.Port,
		ID:               d.ID,
		AddPaused:        d.AddPaused,
		UseSsl:           d.UseSsl,
		Enable:           d.Enable,
	}
}

func (d *DownloadClientTransmission) fromDownloadClient(client *DownloadClient) {
	d.Tags = client.Tags
	d.Name = client.Name
	d.Host = client.Host
	d.URLBase = client.URLBase
	d.Username = client.Username
	d.Password = client.Password
	d.TvCategory = client.TvCategory
	d.TvDirectory = client.TvDirectory
	d.RecentTvPriority = client.RecentTvPriority
	d.OlderTvPriority = client.OlderTvPriority
	d.Priority = client.Priority
	d.Port = client.Port
	d.ID = client.ID
	d.AddPaused = client.AddPaused
	d.UseSsl = client.UseSsl
	d.Enable = client.Enable
}

func (r *DownloadClientTransmissionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + downloadClientTransmissionResourceName
}

func (r *DownloadClientTransmissionResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		MarkdownDescription: "<!-- subcategory:Download Clients -->Download Client Transmission resource.\nFor more information refer to [Download Client](https://wiki.servarr.com/prowlarr/settings#download-clients) and [Transmission](https://wiki.servarr.com/prowlarr/supported#transmission).",
		Attributes: map[string]tfsdk.Attribute{
			"enable": {
				MarkdownDescription: "Enable flag.",
				Optional:            true,
				Computed:            true,
				Type:                types.BoolType,
			},
			"priority": {
				MarkdownDescription: "Priority.",
				Optional:            true,
				Computed:            true,
				Type:                types.Int64Type,
			},
			"name": {
				MarkdownDescription: "Download Client name.",
				Required:            true,
				Type:                types.StringType,
			},
			"tags": {
				MarkdownDescription: "List of associated tags.",
				Optional:            true,
				Computed:            true,
				Type: types.SetType{
					ElemType: types.Int64Type,
				},
			},
			"id": {
				MarkdownDescription: "Download Client ID.",
				Computed:            true,
				Type:                types.Int64Type,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
			},
			// Field values
			"add_paused": {
				MarkdownDescription: "Add paused flag.",
				Optional:            true,
				Computed:            true,
				Type:                types.BoolType,
			},
			"use_ssl": {
				MarkdownDescription: "Use SSL flag.",
				Optional:            true,
				Computed:            true,
				Type:                types.BoolType,
			},
			"port": {
				MarkdownDescription: "Port.",
				Optional:            true,
				Computed:            true,
				Type:                types.Int64Type,
			},
			"recent_tv_priority": {
				MarkdownDescription: "Recent TV priority. `0` Last, `1` First.",
				Optional:            true,
				Computed:            true,
				Type:                types.Int64Type,
				Validators: []tfsdk.AttributeValidator{
					tools.IntMatch([]int64{0, 1}),
				},
			},
			"older_tv_priority": {
				MarkdownDescription: "Older TV priority. `0` Last, `1` First.",
				Optional:            true,
				Computed:            true,
				Type:                types.Int64Type,
				Validators: []tfsdk.AttributeValidator{
					tools.IntMatch([]int64{0, 1}),
				},
			},
			"host": {
				MarkdownDescription: "host.",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"url_base": {
				MarkdownDescription: "Base URL.",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"username": {
				MarkdownDescription: "Username.",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"password": {
				MarkdownDescription: "Password.",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"tv_category": {
				MarkdownDescription: "TV category.",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"tv_directory": {
				MarkdownDescription: "TV directory.",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
		},
	}, nil
}

func (r *DownloadClientTransmissionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*prowlarr.Prowlarr)
	if !ok {
		resp.Diagnostics.AddError(
			tools.UnexpectedResourceConfigureType,
			fmt.Sprintf("Expected *prowlarr.Prowlarr, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *DownloadClientTransmissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var client *DownloadClientTransmission

	resp.Diagnostics.Append(req.Plan.Get(ctx, &client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create new DownloadClientTransmission
	request := client.read(ctx)

	response, err := r.client.AddDownloadClientContext(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError(tools.ClientError, fmt.Sprintf("Unable to create %s, got error: %s", downloadClientTransmissionResourceName, err))

		return
	}

	tflog.Trace(ctx, "created "+downloadClientTransmissionResourceName+": "+strconv.Itoa(int(response.ID)))
	// Generate resource state struct
	client.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &client)...)
}

func (r *DownloadClientTransmissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var client DownloadClientTransmission

	resp.Diagnostics.Append(req.State.Get(ctx, &client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get DownloadClientTransmission current value
	response, err := r.client.GetDownloadClientContext(ctx, client.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(tools.ClientError, fmt.Sprintf("Unable to read %s, got error: %s", downloadClientTransmissionResourceName, err))

		return
	}

	tflog.Trace(ctx, "read "+downloadClientTransmissionResourceName+": "+strconv.Itoa(int(response.ID)))
	// Map response body to resource schema attribute
	client.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &client)...)
}

func (r *DownloadClientTransmissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Get plan values
	var client *DownloadClientTransmission

	resp.Diagnostics.Append(req.Plan.Get(ctx, &client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update DownloadClientTransmission
	request := client.read(ctx)

	response, err := r.client.UpdateDownloadClientContext(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError(tools.ClientError, fmt.Sprintf("Unable to update %s, got error: %s", downloadClientTransmissionResourceName, err))

		return
	}

	tflog.Trace(ctx, "updated "+downloadClientTransmissionResourceName+": "+strconv.Itoa(int(response.ID)))
	// Generate resource state struct
	client.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &client)...)
}

func (r *DownloadClientTransmissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var client *DownloadClientTransmission

	resp.Diagnostics.Append(req.State.Get(ctx, &client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete DownloadClientTransmission current value
	err := r.client.DeleteDownloadClientContext(ctx, client.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(tools.ClientError, fmt.Sprintf("Unable to read %s, got error: %s", downloadClientTransmissionResourceName, err))

		return
	}

	tflog.Trace(ctx, "deleted "+downloadClientTransmissionResourceName+": "+strconv.Itoa(int(client.ID.ValueInt64())))
	resp.State.RemoveResource(ctx)
}

func (r *DownloadClientTransmissionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			tools.UnexpectedImportIdentifier,
			fmt.Sprintf("Expected import identifier with format: ID. Got: %q", req.ID),
		)

		return
	}

	tflog.Trace(ctx, "imported "+downloadClientTransmissionResourceName+": "+strconv.Itoa(id))
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (d *DownloadClientTransmission) write(ctx context.Context, downloadClient *prowlarr.DownloadClientOutput) {
	genericDownloadClient := DownloadClient{
		Enable:   types.BoolValue(downloadClient.Enable),
		Priority: types.Int64Value(int64(downloadClient.Priority)),
		ID:       types.Int64Value(downloadClient.ID),
		Name:     types.StringValue(downloadClient.Name),
		Tags:     types.SetValueMust(types.Int64Type, nil),
	}
	tfsdk.ValueFrom(ctx, downloadClient.Tags, genericDownloadClient.Tags.Type(ctx), &genericDownloadClient.Tags)
	genericDownloadClient.writeFields(ctx, downloadClient.Fields)
	d.fromDownloadClient(&genericDownloadClient)
}

func (d *DownloadClientTransmission) read(ctx context.Context) *prowlarr.DownloadClientInput {
	var tags []int

	tfsdk.ValueAs(ctx, d.Tags, &tags)

	return &prowlarr.DownloadClientInput{
		Enable:         d.Enable.ValueBool(),
		Priority:       int(d.Priority.ValueInt64()),
		ID:             d.ID.ValueInt64(),
		ConfigContract: DownloadClientTransmissionConfigContrat,
		Implementation: DownloadClientTransmissionImplementation,
		Name:           d.Name.ValueString(),
		Protocol:       DownloadClientTransmissionProtocol,
		Tags:           tags,
		Fields:         d.toDownloadClient().readFields(ctx),
	}
}