package service

// Followers
// EntryData.ProfilePage[0].Graphql.User.EdgeFollowedBy.Count

// EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia.Edges[...loop] // up to 12.
// EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia.Edges[...loop].Node.EdgeMediaToComment.Count // up to 12.
// EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia.Edges[...loop].Node.EdgeLikedBy.Count // up to 12.

//

type WebDumpedJson struct {
	CountryCode  string `json:"country_code"`
	LanguageCode string `json:"language_code"`
	Locale       string `json:"locale"`
	EntryData    struct {
		ProfilePage []Profile `json:"ProfilePage"`
	} `json:"entry_data"`
}

type Profile struct {
	SeoCategoryInfos [][]string `json:"seo_category_infos"`
	Graphql          struct {
		User struct {
			Biography              string      `json:"biography"`
			BlockedByViewer        bool        `json:"blocked_by_viewer"`
			RestrictedByViewer     interface{} `json:"restricted_by_viewer"`
			CountryBlock           bool        `json:"country_block"`
			ExternalURL            string      `json:"external_url"`
			ExternalURLLinkshimmed string      `json:"external_url_linkshimmed"`
			EdgeFollowedBy         struct {
				Count int `json:"count"`
			} `json:"edge_followed_by"`
			FollowedByViewer bool `json:"followed_by_viewer"`
			EdgeFollow       struct {
				Count int `json:"count"`
			} `json:"edge_follow"`
			FollowsViewer         bool        `json:"follows_viewer"`
			FullName              string      `json:"full_name"`
			HasArEffects          bool        `json:"has_ar_effects"`
			HasClips              bool        `json:"has_clips"`
			HasGuides             bool        `json:"has_guides"`
			HasChannel            bool        `json:"has_channel"`
			HighlightReelCount    int         `json:"highlight_reel_count"`
			HasRequestedViewer    bool        `json:"has_requested_viewer"`
			HideLikeAndViewCounts bool        `json:"hide_like_and_view_counts"`
			ID                    string      `json:"id"`
			IsBusinessAccount     bool        `json:"is_business_account"`
			IsProfessionalAccount bool        `json:"is_professional_account"`
			IsJoinedRecently      bool        `json:"is_joined_recently"`
			BusinessAddressJSON   interface{} `json:"business_address_json"`
			BusinessContactMethod interface{} `json:"business_contact_method"`
			BusinessEmail         interface{} `json:"business_email"`
			BusinessPhoneNumber   interface{} `json:"business_phone_number"`
			BusinessCategoryName  interface{} `json:"business_category_name"`
			OverallCategoryName   interface{} `json:"overall_category_name"`
			CategoryEnum          interface{} `json:"category_enum"`
			CategoryName          string      `json:"category_name"`
			IsPrivate             bool        `json:"is_private"`
			IsVerified            bool        `json:"is_verified"`
			EdgeMutualFollowedBy  struct {
				Count int           `json:"count"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_mutual_followed_by"`
			ProfilePicURL            string        `json:"profile_pic_url"`
			ProfilePicURLHd          string        `json:"profile_pic_url_hd"`
			RequestedByViewer        bool          `json:"requested_by_viewer"`
			ShouldShowCategory       bool          `json:"should_show_category"`
			ShouldShowPublicContacts bool          `json:"should_show_public_contacts"`
			Username                 string        `json:"username"`
			ConnectedFbPage          interface{}   `json:"connected_fb_page"`
			Pronouns                 []interface{} `json:"pronouns"`
			EdgeFelixVideoTimeline   struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool   `json:"has_next_page"`
					EndCursor   string `json:"end_cursor"`
				} `json:"page_info"`
				Edges []struct {
					Node struct {
						Typename   string `json:"__typename"`
						ID         string `json:"id"`
						Shortcode  string `json:"shortcode"`
						Dimensions struct {
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"dimensions"`
						DisplayURL            string `json:"display_url"`
						EdgeMediaToTaggedUser struct {
							Edges []interface{} `json:"edges"`
						} `json:"edge_media_to_tagged_user"`
						FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
						FactCheckInformation   interface{} `json:"fact_check_information"`
						GatingInfo             interface{} `json:"gating_info"`
						SharingFrictionInfo    struct {
							ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
							BloksAppURL               interface{} `json:"bloks_app_url"`
						} `json:"sharing_friction_info"`
						MediaOverlayInfo interface{} `json:"media_overlay_info"`
						MediaPreview     string      `json:"media_preview"`
						Owner            struct {
							ID       string `json:"id"`
							Username string `json:"username"`
						} `json:"owner"`
						IsVideo              bool        `json:"is_video"`
						HasUpcomingEvent     bool        `json:"has_upcoming_event"`
						AccessibilityCaption interface{} `json:"accessibility_caption"`
						DashInfo             struct {
							IsDashEligible    bool        `json:"is_dash_eligible"`
							VideoDashManifest interface{} `json:"video_dash_manifest"`
							NumberOfQualities int         `json:"number_of_qualities"`
						} `json:"dash_info"`
						HasAudio           bool   `json:"has_audio"`
						TrackingToken      string `json:"tracking_token"`
						VideoURL           string `json:"video_url"`
						VideoViewCount     int    `json:"video_view_count"`
						EdgeMediaToCaption struct {
							Edges []struct {
								Node struct {
									Text string `json:"text"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_media_to_caption"`
						EdgeMediaToComment struct {
							Count int `json:"count"`
						} `json:"edge_media_to_comment"`
						CommentsDisabled bool `json:"comments_disabled"`
						TakenAtTimestamp int  `json:"taken_at_timestamp"`
						EdgeLikedBy      struct {
							Count int `json:"count"`
						} `json:"edge_liked_by"`
						EdgeMediaPreviewLike struct {
							Count int `json:"count"`
						} `json:"edge_media_preview_like"`
						Location           interface{} `json:"location"`
						ThumbnailSrc       string      `json:"thumbnail_src"`
						ThumbnailResources []struct {
							Src          string `json:"src"`
							ConfigWidth  int    `json:"config_width"`
							ConfigHeight int    `json:"config_height"`
						} `json:"thumbnail_resources"`
						FelixProfileGridCrop interface{}   `json:"felix_profile_grid_crop"`
						CoauthorProducers    []interface{} `json:"coauthor_producers"`
						EncodingStatus       interface{}   `json:"encoding_status"`
						IsPublished          bool          `json:"is_published"`
						ProductType          string        `json:"product_type"`
						Title                string        `json:"title"`
						VideoDuration        float64       `json:"video_duration"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_felix_video_timeline"`
			EdgeOwnerToTimelineMedia struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool   `json:"has_next_page"`
					EndCursor   string `json:"end_cursor"`
				} `json:"page_info"`
				Edges []Edge `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
			EdgeMediaCollections struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool        `json:"has_next_page"`
					EndCursor   interface{} `json:"end_cursor"`
				} `json:"page_info"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_collections"`
		} `json:"user"`
	} `json:"graphql"`
	ToastContentOnLoad      interface{} `json:"toast_content_on_load"`
	ShowViewShop            bool        `json:"show_view_shop"`
	ProfilePicEditSyncProps struct {
		ShowChangeProfilePicConfirmDialog bool   `json:"show_change_profile_pic_confirm_dialog"`
		ShowProfilePicSyncReminders       bool   `json:"show_profile_pic_sync_reminders"`
		IdentityID                        string `json:"identity_id"`
		IsBusinessCentralIdentity         bool   `json:"is_business_central_identity"`
	} `json:"profile_pic_edit_sync_props"`
}

type Node struct {
	Typename              string     `json:"__typename"`
	ID                    string     `json:"id"`
	Shortcode             string     `json:"shortcode"`
	Dimensions            Dimensions `json:"dimensions"`
	DisplayURL            string     `json:"display_url"`
	EdgeMediaToTaggedUser struct {
		Edges []interface{} `json:"edges"`
	} `json:"edge_media_to_tagged_user"`
	FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
	FactCheckInformation   interface{} `json:"fact_check_information"`
	GatingInfo             interface{} `json:"gating_info"`
	SharingFrictionInfo    struct {
		ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
		BloksAppURL               interface{} `json:"bloks_app_url"`
	} `json:"sharing_friction_info"`
	MediaOverlayInfo interface{} `json:"media_overlay_info"`
	MediaPreview     string      `json:"media_preview"`
	Owner            struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	} `json:"owner"`
	IsVideo              bool   `json:"is_video"`
	HasUpcomingEvent     bool   `json:"has_upcoming_event"`
	AccessibilityCaption string `json:"accessibility_caption"`
	EdgeMediaToCaption   struct {
		Edges []struct {
			Node struct {
				Text string `json:"text"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"edge_media_to_caption"`
	EdgeMediaToComment struct {
		Count int `json:"count"`
	} `json:"edge_media_to_comment"`
	CommentsDisabled bool `json:"comments_disabled"`
	TakenAtTimestamp int  `json:"taken_at_timestamp"`
	EdgeLikedBy      struct {
		Count int `json:"count"`
	} `json:"edge_liked_by"`
	EdgeMediaPreviewLike struct {
		Count int `json:"count"`
	} `json:"edge_media_preview_like"`
	Location struct {
		ID            string `json:"id"`
		HasPublicPage bool   `json:"has_public_page"`
		Name          string `json:"name"`
		Slug          string `json:"slug"`
	} `json:"location"`
	ThumbnailSrc       string `json:"thumbnail_src"`
	ThumbnailResources []struct {
		Src          string `json:"src"`
		ConfigWidth  int    `json:"config_width"`
		ConfigHeight int    `json:"config_height"`
	} `json:"thumbnail_resources"`
	CoauthorProducers []interface{} `json:"coauthor_producers"`
}

type Dimensions struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type Edge struct {
	Node Node `json:"node,omitempty"`
}
