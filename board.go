package trello

type membersInvited struct {
	baseMember
	url   string
	Prefs staticField
}

func createMembersInvited(m model) membersInvited {
	mm := membersInvited{
		url: m.getURL() + "/membersInvited",
	}
	memb := createBaseMember(m)
	memInv := membersInvited{
		memb,
		mm.url,
		staticField(mm.url + "/prefs"),
	}
	memInv.Actions = staticField(mm.url + "/actions")
	memInv.Avatar = staticField(mm.url + "/avatar")
	memInv.AvatarHash = staticField(mm.url + "/avatarHash")
	memInv.AvatarSource = staticField(mm.url + "/avatarSource")
	memInv.Bio = staticField(mm.url + "/bio")
	memInv.BioData = staticField(mm.url + "/bioData")
	memInv.Confirmed = staticField(mm.url + "/confirmed")
	memInv.Deltas = staticField(mm.url + "/deltas")
	memInv.Email = staticField(mm.url + "/email")
	memInv.FullName = staticField(mm.url + "/fullName")
	memInv.GravatarHash = staticField(mm.url + "/gravatarHash")
	memInv.IdBoards = staticField(mm.url + "/idBoards")
	memInv.IdBoardsPinned = staticField(mm.url + "/idBoardsPinned")
	memInv.IdOrganizations = staticField(mm.url + "/idOrganizations")
	memInv.IdPremOrgsAdmin = staticField(mm.url + "/idPremOrgsAdmin")
	memInv.Initials = staticField(mm.url + "/initials")
	memInv.LoginTypes = staticField(mm.url + "/loginTypes")
	memInv.MemberType = staticField(mm.url + "/memberType")
	memInv.OneTimeMessagesDismissed = staticField(mm.url + "/oneTimeMessagesDismissed")
	memInv.PremiumFeatures = staticField(mm.url + "/premiumFeatures")
	memInv.Products = staticField(mm.url + "/products")
	memInv.Status = staticField(mm.url + "/status")
	memInv.Tokens = staticField(mm.url + "/tokens")
	memInv.Trophies = staticField(mm.url + "/trophies")
	memInv.UploadedAvatarHash = staticField(mm.url + "/uploadedAvatarHash")
	memInv.Url = staticField(mm.url + "/url")
	memInv.Username = staticField(mm.url + "/username")
	return memInv
}

func (m membersInvited) getURL() string {
	return m.url
}

type memberships struct {
	url string
}

func createMemberships(m model) memberships {
	return memberships{
		url: m.getURL() + "/memberships",
	}
}

func (m memberships) ID(id string) memberships {
	m.url = m.url + "/" + id
	return m
}

func (m memberships) getURL() string {
	return m.url
}

type myPrefs struct {
	EmailPosition,
	IdEmailList,
	ShowListGuide,
	ShowSidebar,
	ShowSidebarActivity,
	ShowSidebarBoardActions,
	ShowSidebarMembers staticField
}

func createMyPrefs(m model) myPrefs {
	mpURL := m.getURL() + "/myPrefs"
	return myPrefs{
		EmailPosition:           staticField(mpURL + "/emailPosition"),
		IdEmailList:             staticField(mpURL + "/idEmailList"),
		ShowListGuide:           staticField(mpURL + "/showListGuide"),
		ShowSidebar:             staticField(mpURL + "/showSidebar"),
		ShowSidebarActivity:     staticField(mpURL + "/showSidebarActivity"),
		ShowSidebarBoardActions: staticField(mpURL + "/showSidebarBoardActions"),
		ShowSidebarMembers:      staticField(mpURL + "/showSidebarMembers"),
	}
}

type boardPrefs struct {
	url string
	Background,
	CalendarFeedEnabled,
	CardAging,
	CardCovers,
	Comments,
	Invitations,
	PermissionLevel,
	SelfJoin,
	Voting staticField
}

func createBoardPrefs(m model) boardPrefs {
	pURL := m.getURL() + "/prefs"
	return boardPrefs{
		url:                 pURL,
		Background:          staticField(pURL + "/background"),
		CalendarFeedEnabled: staticField(pURL + "/calendarFeedEnabled"),
		CardAging:           staticField(pURL + "/cardAging"),
		CardCovers:          staticField(pURL + "/cardCovers"),
		Comments:            staticField(pURL + "/comments"),
		Invitations:         staticField(pURL + "/invitations"),
		PermissionLevel:     staticField(pURL + "/permissionLevel"),
		SelfJoin:            staticField(pURL + "/selfJoin"),
		Voting:              staticField(pURL + "/voting"),
	}
}

func (b boardPrefs) getURL() string {
	return b.url
}

type powerUps struct {
	url string
	Calendar,
	CardAging,
	Recap,
	Voting staticField
}

func createPowerUps(m model) powerUps {
	pURL := m.getURL() + "/powerUps"
	return powerUps{
		url:       pURL,
		Calendar:  staticField(pURL + "/calendar"),
		CardAging: staticField(pURL + "/cardAging"),
		Recap:     staticField(pURL + "/recap"),
		Voting:    staticField(pURL + "/voting"),
	}
}

func (p powerUps) getURL() string {
	return p.url
}

type filterBoards struct {
	url string
	All,
	Closed,
	Members,
	Open,
	Organization,
	Pinned,
	Public,
	Starred,
	Unpinned staticField
}

func createFilterBoard(m model) filterBoards {
	bURL := m.getURL() + "/boards"
	return filterBoards{
		url:          bURL,
		All:          staticField(bURL + "/all"),
		Closed:       staticField(bURL + "/closed"),
		Members:      staticField(bURL + "/members"),
		Open:         staticField(bURL + "/open"),
		Organization: staticField(bURL + "/organization"),
		Pinned:       staticField(bURL + "/pinned"),
		Public:       staticField(bURL + "/public"),
		Starred:      staticField(bURL + "/starred"),
		Unpinned:     staticField(bURL + "/unpinned"),
	}
}

func (b filterBoards) getURL() string {
	return b.url
}

type boardStars struct {
	url string
	IdBoard,
	Pos staticField
}

func createBoardStars(m model) boardStars {
	return boardStars{
		url: m.getURL() + "/boardStars",
	}
}

func (b boardStars) ID(id string) boardStars {
	bURL := b.url + "/" + id

	return boardStars{
		url:     bURL,
		IdBoard: staticField(bURL + "/idBoard"),
		Pos:     staticField(bURL + "/pos"),
	}
}

func (b boardStars) getURL() string {
	return b.url
}

type baseBoard struct {
	url string
	//static fields
	Actions,
	BoardStars,
	Checklists,
	Closed,
	DateLastActivity,
	DateLastView,
	Deltas,
	Desc,
	DescData,
	IdOrganization,
	Invitations,
	Invited,
	MarkAsViewed,
	Name,
	Pinned,
	ShortLink,
	ShortUrl,
	Starred,
	Subscribed,
	Url staticField
	//submodels
	LabelNames  labelNames
	Memberships memberships
	PowerUps    powerUps
	Prefs       boardPrefs
}

func createBaseBoard(m model) baseBoard {
	bURL := m.getURL() + "/board"
	b := baseBoard{}

	b.url = bURL
	b.Actions = staticField(bURL + "/actions")
	b.BoardStars = staticField(bURL + "/boardStars")
	b.Checklists = staticField(bURL + "/checklists")
	b.Closed = staticField(bURL + "/closed")
	b.DateLastActivity = staticField(bURL + "/dateLastActivity")
	b.DateLastView = staticField(bURL + "/dateLastView")
	b.Deltas = staticField(bURL + "/deltas")
	b.Desc = staticField(bURL + "/desc")
	b.DescData = staticField(bURL + "/descData")
	b.IdOrganization = staticField(bURL + "/idOrganization")
	b.Invitations = staticField(bURL + "/invitations")
	b.Invited = staticField(bURL + "/invited")
	b.MarkAsViewed = staticField(bURL + "/markAsViewed")
	b.Name = staticField(bURL + "/name")
	b.Pinned = staticField(bURL + "/pinned")
	b.ShortLink = staticField(bURL + "/shortLink")
	b.ShortUrl = staticField(bURL + "/shortUrl")
	b.Starred = staticField(bURL + "/starred")
	b.Subscribed = staticField(bURL + "/subscribed")
	b.Url = staticField(bURL + "/url")
	b.LabelNames = createLabelNames(b)
	b.Memberships = createMemberships(b)
	b.PowerUps = createPowerUps(b)
	b.Prefs = createBoardPrefs(b)
	return b
}

func (b baseBoard) getURL() string {
	return b.url
}

type board struct {
	baseBoard
	url string
	//submodels
	CalendarKey    calKey
	EmailKey       emailKey
	Cards          filterCard
	Labels         label
	Lists          list
	Members        member
	MembersInvited membersInvited

	MyPrefs      myPrefs
	Organization organization
}

//Boards is the struct representing the boards model from trello
var Boards = board{
	url: "/boards",
}

func (b board) ID(id string) board {
	boardURL := b.url + "/" + id
	b.url = boardURL
	b.Actions = staticField(boardURL + "/actions")
	b.BoardStars = staticField(boardURL + "/boardStars")
	b.Checklists = staticField(boardURL + "/checklists")
	b.Closed = staticField(boardURL + "/closed")
	b.DateLastActivity = staticField(boardURL + "/dateLastActivity")
	b.DateLastView = staticField(boardURL + "/dateLastView")
	b.Deltas = staticField(boardURL + "/deltas")
	b.Desc = staticField(boardURL + "/desc")
	b.DescData = staticField(boardURL + "/descData")
	b.IdOrganization = staticField(boardURL + "/idOrganization")
	b.Invitations = staticField(boardURL + "/invitations")
	b.Invited = staticField(boardURL + "/invited")
	b.Name = staticField(boardURL + "/name")
	b.Pinned = staticField(boardURL + "/pinned")
	b.Prefs = createBoardPrefs(b)
	b.PowerUps = createPowerUps(b)
	b.ShortLink = staticField(boardURL + "/shortLink")
	b.ShortUrl = staticField(boardURL + "/shortUrl")
	b.Starred = staticField(boardURL + "/starred")
	b.Subscribed = staticField(boardURL + "/subscribed")
	b.Url = staticField(boardURL + "/url")
	b.Cards = createFilterCard(b)
	b.CalendarKey = createCalKey(b)
	b.EmailKey = createEmailKey(b)
	b.LabelNames = createLabelNames(b)
	b.Labels = createLabel(b)
	b.Lists = createList(b)
	b.MarkAsViewed = staticField(boardURL + "/markAsViewed")
	b.Members = createMember(b)
	b.MembersInvited = createMembersInvited(b)
	b.Memberships = createMemberships(b)
	b.MyPrefs = createMyPrefs(b)
	b.Organization = createOrganization(b)
	return b
}

func (b board) getURL() string {
	return b.url
}
