package trello

import "testing"

func testNotifications(t *testing.T) {

	testModels := []model{
		Notifications.ID(ID),
		Notifications.ID(ID).Data,
		Notifications.ID(ID).Date,
		Notifications.ID(ID).IdMemberCreator,
		Notifications.ID(ID).Type,
		Notifications.ID(ID).Unread,
		Notifications.ID(ID).Board,
		Notifications.ID(ID).Board.Closed,
		Notifications.ID(ID).Board.DateLastActivity,
		Notifications.ID(ID).Board.DateLastView,
		Notifications.ID(ID).Board.Desc,
		Notifications.ID(ID).Board.DescData,
		Notifications.ID(ID).Board.IdOrganization,
		Notifications.ID(ID).Board.Invitations,
		Notifications.ID(ID).Board.Invited,
		Notifications.ID(ID).Board.LabelNames,
		Notifications.ID(ID).Board.Memberships,
		Notifications.ID(ID).Board.Name,
		Notifications.ID(ID).Board.Pinned,
		Notifications.ID(ID).Board.PowerUps,
		Notifications.ID(ID).Board.Prefs,
		Notifications.ID(ID).Board.ShortLink,
		Notifications.ID(ID).Board.ShortUrl,
		Notifications.ID(ID).Board.Starred,
		Notifications.ID(ID).Board.Subscribed,
		Notifications.ID(ID).Board.Url,
		Notifications.ID(ID).Card,
		Notifications.ID(ID).Card.Badges,
		Notifications.ID(ID).Card.CheckItemStates,
		Notifications.ID(ID).Card.Closed,
		Notifications.ID(ID).Card.DateLastActivity,
		Notifications.ID(ID).Card.Desc,
		Notifications.ID(ID).Card.DescData,
		Notifications.ID(ID).Card.Due,
		Notifications.ID(ID).Card.Email,
		Notifications.ID(ID).Card.IdAttachmentCover,
		Notifications.ID(ID).Card.IdBoard,
		Notifications.ID(ID).Card.IdChecklists,
		Notifications.ID(ID).Card.IdLabels,
		Notifications.ID(ID).Card.IdList,
		Notifications.ID(ID).Card.IdMembers,
		Notifications.ID(ID).Card.IdMembersVoted,
		Notifications.ID(ID).Card.IdShort,
		Notifications.ID(ID).Card.Labels,
		Notifications.ID(ID).Card.ManualCoverAttachment,
		Notifications.ID(ID).Card.Name,
		Notifications.ID(ID).Card.Pos,
		Notifications.ID(ID).Card.ShortLink,
		Notifications.ID(ID).Card.ShortUrl,
		Notifications.ID(ID).Card.Subscribed,
		Notifications.ID(ID).Card.Url,
		Notifications.ID(ID).Display,
		Notifications.ID(ID).Entities,
		Notifications.ID(ID).List,
		Notifications.ID(ID).List.Closed,
		Notifications.ID(ID).List.IdBoard,
		Notifications.ID(ID).List.Name,
		Notifications.ID(ID).List.Pos,
		Notifications.ID(ID).List.Subscribed,
		Notifications.ID(ID).Member,
		Notifications.ID(ID).Member.AvatarHash,
		Notifications.ID(ID).Member.AvatarSource,
		Notifications.ID(ID).Member.Bio,
		Notifications.ID(ID).Member.BioData,
		Notifications.ID(ID).Member.Confirmed,
		Notifications.ID(ID).Member.Email,
		Notifications.ID(ID).Member.FullName,
		Notifications.ID(ID).Member.GravatarHash,
		Notifications.ID(ID).Member.IdBoards,
		Notifications.ID(ID).Member.IdBoardsPinned,
		Notifications.ID(ID).Member.IdOrganizations,
		Notifications.ID(ID).Member.IdPremOrgsAdmin,
		Notifications.ID(ID).Member.Initials,
		Notifications.ID(ID).Member.LoginTypes,
		Notifications.ID(ID).Member.MemberType,
		Notifications.ID(ID).Member.OneTimeMessagesDismissed,
		Notifications.ID(ID).Member.Prefs,
		Notifications.ID(ID).Member.PremiumFeatures,
		Notifications.ID(ID).Member.Products,
		Notifications.ID(ID).Member.Status,
		Notifications.ID(ID).Member.Trophies,
		Notifications.ID(ID).Member.UploadedAvatarHash,
		Notifications.ID(ID).Member.Url,
		Notifications.ID(ID).Member.Username,
		Notifications.ID(ID).MemberCreator,
		Notifications.ID(ID).MemberCreator.AvatarHash,
		Notifications.ID(ID).MemberCreator.AvatarSource,
		Notifications.ID(ID).MemberCreator.Bio,
		Notifications.ID(ID).MemberCreator.BioData,
		Notifications.ID(ID).MemberCreator.Confirmed,
		Notifications.ID(ID).MemberCreator.Email,
		Notifications.ID(ID).MemberCreator.FullName,
		Notifications.ID(ID).MemberCreator.GravatarHash,
		Notifications.ID(ID).MemberCreator.IdBoards,
		Notifications.ID(ID).MemberCreator.IdBoardsPinned,
		Notifications.ID(ID).MemberCreator.IdOrganizations,
		Notifications.ID(ID).MemberCreator.IdPremOrgsAdmin,
		Notifications.ID(ID).MemberCreator.Initials,
		Notifications.ID(ID).MemberCreator.LoginTypes,
		Notifications.ID(ID).MemberCreator.MemberType,
		Notifications.ID(ID).MemberCreator.OneTimeMessagesDismissed,
		Notifications.ID(ID).MemberCreator.Prefs,
		Notifications.ID(ID).MemberCreator.PremiumFeatures,
		Notifications.ID(ID).MemberCreator.Products,
		Notifications.ID(ID).MemberCreator.Status,
		Notifications.ID(ID).MemberCreator.Trophies,
		Notifications.ID(ID).MemberCreator.UploadedAvatarHash,
		Notifications.ID(ID).MemberCreator.Url,
		Notifications.ID(ID).MemberCreator.Username,
		Notifications.ID(ID).Organization,
		Notifications.ID(ID).Organization.BillableMemberCount,
		Notifications.ID(ID).Organization.Desc,
		Notifications.ID(ID).Organization.DescData,
		Notifications.ID(ID).Organization.DisplayName,
		Notifications.ID(ID).Organization.IdBoards,
		Notifications.ID(ID).Organization.Invitations,
		Notifications.ID(ID).Organization.Invited,
		Notifications.ID(ID).Organization.LogoHash,
		Notifications.ID(ID).Organization.Memberships,
		Notifications.ID(ID).Organization.Name,
		Notifications.ID(ID).Organization.PowerUps,
		Notifications.ID(ID).Organization.Prefs,
		Notifications.ID(ID).Organization.PremiumFeatures,
		Notifications.ID(ID).Organization.Products,
		Notifications.ID(ID).Organization.Url,
		Notifications.ID(ID).Organization.Website,
		Notifications.All.Read,
	}

	testExpects := []string{
		"/notifications/" + ID,
		"/notifications/" + ID + "/data",
		"/notifications/" + ID + "/date",
		"/notifications/" + ID + "/idMemberCreator",
		"/notifications/" + ID + "/type",
		"/notifications/" + ID + "/unread",
		"/notifications/" + ID + "/board",
		"/notifications/" + ID + "/board/closed",
		"/notifications/" + ID + "/board/dateLastActivity",
		"/notifications/" + ID + "/board/dateLastView",
		"/notifications/" + ID + "/board/desc",
		"/notifications/" + ID + "/board/descData",
		"/notifications/" + ID + "/board/idOrganization",
		"/notifications/" + ID + "/board/invitations",
		"/notifications/" + ID + "/board/invited",
		"/notifications/" + ID + "/board/labelNames",
		"/notifications/" + ID + "/board/memberships",
		"/notifications/" + ID + "/board/name",
		"/notifications/" + ID + "/board/pinned",
		"/notifications/" + ID + "/board/powerUps",
		"/notifications/" + ID + "/board/prefs",
		"/notifications/" + ID + "/board/shortLink",
		"/notifications/" + ID + "/board/shortUrl",
		"/notifications/" + ID + "/board/starred",
		"/notifications/" + ID + "/board/subscribed",
		"/notifications/" + ID + "/board/url",
		"/notifications/" + ID + "/card",
		"/notifications/" + ID + "/card/badges",
		"/notifications/" + ID + "/card/checkItemStates",
		"/notifications/" + ID + "/card/closed",
		"/notifications/" + ID + "/card/dateLastActivity",
		"/notifications/" + ID + "/card/desc",
		"/notifications/" + ID + "/card/descData",
		"/notifications/" + ID + "/card/due",
		"/notifications/" + ID + "/card/email",
		"/notifications/" + ID + "/card/idAttachmentCover",
		"/notifications/" + ID + "/card/idBoard",
		"/notifications/" + ID + "/card/idChecklists",
		"/notifications/" + ID + "/card/idLabels",
		"/notifications/" + ID + "/card/idList",
		"/notifications/" + ID + "/card/idMembers",
		"/notifications/" + ID + "/card/idMembersVoted",
		"/notifications/" + ID + "/card/idShort",
		"/notifications/" + ID + "/card/labels",
		"/notifications/" + ID + "/card/manualCoverAttachment",
		"/notifications/" + ID + "/card/name",
		"/notifications/" + ID + "/card/pos",
		"/notifications/" + ID + "/card/shortLink",
		"/notifications/" + ID + "/card/shortUrl",
		"/notifications/" + ID + "/card/subscribed",
		"/notifications/" + ID + "/card/url",
		"/notifications/" + ID + "/display",
		"/notifications/" + ID + "/entities",
		"/notifications/" + ID + "/list",
		"/notifications/" + ID + "/list/closed",
		"/notifications/" + ID + "/list/idBoard",
		"/notifications/" + ID + "/list/name",
		"/notifications/" + ID + "/list/pos",
		"/notifications/" + ID + "/list/subscribed",
		"/notifications/" + ID + "/member",
		"/notifications/" + ID + "/member/avatarHash",
		"/notifications/" + ID + "/member/avatarSource",
		"/notifications/" + ID + "/member/bio",
		"/notifications/" + ID + "/member/bioData",
		"/notifications/" + ID + "/member/confirmed",
		"/notifications/" + ID + "/member/email",
		"/notifications/" + ID + "/member/fullName",
		"/notifications/" + ID + "/member/gravatarHash",
		"/notifications/" + ID + "/member/idBoards",
		"/notifications/" + ID + "/member/idBoardsPinned",
		"/notifications/" + ID + "/member/idOrganizations",
		"/notifications/" + ID + "/member/idPremOrgsAdmin",
		"/notifications/" + ID + "/member/initials",
		"/notifications/" + ID + "/member/loginTypes",
		"/notifications/" + ID + "/member/memberType",
		"/notifications/" + ID + "/member/oneTimeMessagesDismissed",
		"/notifications/" + ID + "/member/prefs",
		"/notifications/" + ID + "/member/premiumFeatures",
		"/notifications/" + ID + "/member/products",
		"/notifications/" + ID + "/member/status",
		"/notifications/" + ID + "/member/trophies",
		"/notifications/" + ID + "/member/uploadedAvatarHash",
		"/notifications/" + ID + "/member/url",
		"/notifications/" + ID + "/member/username",
		"/notifications/" + ID + "/memberCreator",
		"/notifications/" + ID + "/memberCreator/avatarHash",
		"/notifications/" + ID + "/memberCreator/avatarSource",
		"/notifications/" + ID + "/memberCreator/bio",
		"/notifications/" + ID + "/memberCreator/bioData",
		"/notifications/" + ID + "/memberCreator/confirmed",
		"/notifications/" + ID + "/memberCreator/email",
		"/notifications/" + ID + "/memberCreator/fullName",
		"/notifications/" + ID + "/memberCreator/gravatarHash",
		"/notifications/" + ID + "/memberCreator/idBoards",
		"/notifications/" + ID + "/memberCreator/idBoardsPinned",
		"/notifications/" + ID + "/memberCreator/idOrganizations",
		"/notifications/" + ID + "/memberCreator/idPremOrgsAdmin",
		"/notifications/" + ID + "/memberCreator/initials",
		"/notifications/" + ID + "/memberCreator/loginTypes",
		"/notifications/" + ID + "/memberCreator/memberType",
		"/notifications/" + ID + "/memberCreator/oneTimeMessagesDismissed",
		"/notifications/" + ID + "/memberCreator/prefs",
		"/notifications/" + ID + "/memberCreator/premiumFeatures",
		"/notifications/" + ID + "/memberCreator/products",
		"/notifications/" + ID + "/memberCreator/status",
		"/notifications/" + ID + "/memberCreator/trophies",
		"/notifications/" + ID + "/memberCreator/uploadedAvatarHash",
		"/notifications/" + ID + "/memberCreator/url",
		"/notifications/" + ID + "/memberCreator/username",
		"/notifications/" + ID + "/organization",
		"/notifications/" + ID + "/organization/billableMemberCount",
		"/notifications/" + ID + "/organization/desc",
		"/notifications/" + ID + "/organization/descData",
		"/notifications/" + ID + "/organization/displayName",
		"/notifications/" + ID + "/organization/idBoards",
		"/notifications/" + ID + "/organization/invitations",
		"/notifications/" + ID + "/organization/invited",
		"/notifications/" + ID + "/organization/logoHash",
		"/notifications/" + ID + "/organization/memberships",
		"/notifications/" + ID + "/organization/name",
		"/notifications/" + ID + "/organization/powerUps",
		"/notifications/" + ID + "/organization/prefs",
		"/notifications/" + ID + "/organization/premiumFeatures",
		"/notifications/" + ID + "/organization/products",
		"/notifications/" + ID + "/organization/url",
		"/notifications/" + ID + "/organization/website",
		"/notifications/all/read",
	}
	for i := range testModels {
		m := testModels[i]
		e := testExpects[i]
		if s := GetURL(m); s != e {
			t.Errorf("test model# %d: Expected \"%s\", Got \"%s\"\n", i, e, s)
		}
	}
}
