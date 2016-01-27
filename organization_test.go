package trello

import "testing"

func testOrganizations(t *testing.T) {

 testModels := []model{
  Organizations.ID(ID),
Organizations.ID(ID).BillableMemberCount,
Organizations.ID(ID).Desc,
Organizations.ID(ID).DescData,
Organizations.ID(ID).DisplayName,
Organizations.ID(ID).IdBoards,
Organizations.ID(ID).Invitations,
Organizations.ID(ID).Invited,
Organizations.ID(ID).LogoHash,
Organizations.ID(ID).Memberships,
Organizations.ID(ID).Name,
Organizations.ID(ID).PowerUps,
Organizations.ID(ID).Prefs,
Organizations.ID(ID).PremiumFeatures,
Organizations.ID(ID).Products,
Organizations.ID(ID).Url,
Organizations.ID(ID).Website,
Organizations.ID(ID).Actions,
Organizations.ID(ID).Boards,
Organizations.ID(ID).Boards.All,
Organizations.ID(ID).Boards.Closed,
Organizations.ID(ID).Boards.Members,
Organizations.ID(ID).Boards.Open,
Organizations.ID(ID).Boards.Organization,
Organizations.ID(ID).Boards.Pinned,
Organizations.ID(ID).Boards.Public,
Organizations.ID(ID).Boards.Starred,
Organizations.ID(ID).Boards.Unpinned,
Organizations.ID(ID).Deltas,
Organizations.ID(ID).Members,
Organizations.ID(ID).Members.Admins,
Organizations.ID(ID).Members.All,
Organizations.ID(ID).Members.None,
Organizations.ID(ID).Members.Normal,
Organizations.ID(ID).Members.Owners,
Organizations.ID(ID).Members.ID(ID).Cards,
Organizations.ID(ID).MembersInvited,
Organizations.ID(ID).MembersInvited.AvatarHash,
Organizations.ID(ID).MembersInvited.AvatarSource,
Organizations.ID(ID).MembersInvited.Bio,
Organizations.ID(ID).MembersInvited.BioData,
Organizations.ID(ID).MembersInvited.Confirmed,
Organizations.ID(ID).MembersInvited.Email,
Organizations.ID(ID).MembersInvited.FullName,
Organizations.ID(ID).MembersInvited.GravatarHash,
Organizations.ID(ID).MembersInvited.IdBoards,
Organizations.ID(ID).MembersInvited.IdBoardsPinned,
Organizations.ID(ID).MembersInvited.IdOrganizations,
Organizations.ID(ID).MembersInvited.IdPremOrgsAdmin,
Organizations.ID(ID).MembersInvited.Initials,
Organizations.ID(ID).MembersInvited.LoginTypes,
Organizations.ID(ID).MembersInvited.MemberType,
Organizations.ID(ID).MembersInvited.OneTimeMessagesDismissed,
Organizations.ID(ID).MembersInvited.Prefs,
Organizations.ID(ID).MembersInvited.PremiumFeatures,
Organizations.ID(ID).MembersInvited.Products,
Organizations.ID(ID).MembersInvited.Status,
Organizations.ID(ID).MembersInvited.Trophies,
Organizations.ID(ID).MembersInvited.UploadedAvatarHash,
Organizations.ID(ID).MembersInvited.Url,
Organizations.ID(ID).MembersInvited.Username,
Organizations.ID(ID).Memberships.ID(ID),
Organizations.ID(ID).Members.ID(ID),
Organizations.ID(ID).Members.ID(ID).Deactivated,
Organizations.ID(ID).Prefs.AssociatedDomain,
Organizations.ID(ID).Prefs.BoardVisibilityRestrict.Org,
Organizations.ID(ID).Prefs.BoardVisibilityRestrict.Private,
Organizations.ID(ID).Prefs.BoardVisibilityRestrict.Public,
Organizations.ID(ID).Prefs.ExternalMembersDisabled,
Organizations.ID(ID).Prefs.GoogleAppsVersion,
Organizations.ID(ID).Prefs.OrgInviteRestrict,
Organizations.ID(ID).Prefs.PermissionLevel,
Organizations,
Organizations.ID(ID).Logo,
Organizations.ID(ID).Members.ID(ID).All,

 }

 testExpects := []string{
  "/organizations/"+ID,
"/organizations/"+ID+"/billableMemberCount",
"/organizations/"+ID+"/desc",
"/organizations/"+ID+"/descData",
"/organizations/"+ID+"/displayName",
"/organizations/"+ID+"/idBoards",
"/organizations/"+ID+"/invitations",
"/organizations/"+ID+"/invited",
"/organizations/"+ID+"/logoHash",
"/organizations/"+ID+"/memberships",
"/organizations/"+ID+"/name",
"/organizations/"+ID+"/powerUps",
"/organizations/"+ID+"/prefs",
"/organizations/"+ID+"/premiumFeatures",
"/organizations/"+ID+"/products",
"/organizations/"+ID+"/url",
"/organizations/"+ID+"/website",
"/organizations/"+ID+"/actions",
"/organizations/"+ID+"/boards",
"/organizations/"+ID+"/boards/all",
"/organizations/"+ID+"/boards/closed",
"/organizations/"+ID+"/boards/members",
"/organizations/"+ID+"/boards/open",
"/organizations/"+ID+"/boards/organization",
"/organizations/"+ID+"/boards/pinned",
"/organizations/"+ID+"/boards/public",
"/organizations/"+ID+"/boards/starred",
"/organizations/"+ID+"/boards/unpinned",
"/organizations/"+ID+"/deltas",
"/organizations/"+ID+"/members",
"/organizations/"+ID+"/members/admins",
"/organizations/"+ID+"/members/all",
"/organizations/"+ID+"/members/none",
"/organizations/"+ID+"/members/normal",
"/organizations/"+ID+"/members/owners",
"/organizations/"+ID+"/members/"+ID+"/cards",
"/organizations/"+ID+"/membersInvited",
"/organizations/"+ID+"/membersInvited/avatarHash",
"/organizations/"+ID+"/membersInvited/avatarSource",
"/organizations/"+ID+"/membersInvited/bio",
"/organizations/"+ID+"/membersInvited/bioData",
"/organizations/"+ID+"/membersInvited/confirmed",
"/organizations/"+ID+"/membersInvited/email",
"/organizations/"+ID+"/membersInvited/fullName",
"/organizations/"+ID+"/membersInvited/gravatarHash",
"/organizations/"+ID+"/membersInvited/idBoards",
"/organizations/"+ID+"/membersInvited/idBoardsPinned",
"/organizations/"+ID+"/membersInvited/idOrganizations",
"/organizations/"+ID+"/membersInvited/idPremOrgsAdmin",
"/organizations/"+ID+"/membersInvited/initials",
"/organizations/"+ID+"/membersInvited/loginTypes",
"/organizations/"+ID+"/membersInvited/memberType",
"/organizations/"+ID+"/membersInvited/oneTimeMessagesDismissed",
"/organizations/"+ID+"/membersInvited/prefs",
"/organizations/"+ID+"/membersInvited/premiumFeatures",
"/organizations/"+ID+"/membersInvited/products",
"/organizations/"+ID+"/membersInvited/status",
"/organizations/"+ID+"/membersInvited/trophies",
"/organizations/"+ID+"/membersInvited/uploadedAvatarHash",
"/organizations/"+ID+"/membersInvited/url",
"/organizations/"+ID+"/membersInvited/username",
"/organizations/"+ID+"/memberships/"+ID,
"/organizations/"+ID+"/members/"+ID,
"/organizations/"+ID+"/members/"+ID+"/deactivated",
"/organizations/"+ID+"/prefs/associatedDomain",
"/organizations/"+ID+"/prefs/boardVisibilityRestrict/org",
"/organizations/"+ID+"/prefs/boardVisibilityRestrict/private",
"/organizations/"+ID+"/prefs/boardVisibilityRestrict/public",
"/organizations/"+ID+"/prefs/externalMembersDisabled",
"/organizations/"+ID+"/prefs/googleAppsVersion",
"/organizations/"+ID+"/prefs/orgInviteRestrict",
"/organizations/"+ID+"/prefs/permissionLevel",
"/organizations",
"/organizations/"+ID+"/logo",
"/organizations/"+ID+"/members/"+ID+"/all",
 }
 for i := range testModels {
		m := testModels[i]
		e := testExpects[i]
		if s := GetURL(m); s != e {
			t.Errorf("test model# %d: Expected \"%s\", Got \"%s\"\n", i, e, s)
		}
	}
}
