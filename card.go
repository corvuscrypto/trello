package trello

type attachment struct {
	url string
}

func createAttachment(m model) attachment {
	return attachment{
		url: m.getURL() + "/attachments",
	}
}

func (a attachment) ID(id string) attachment {
	a.url = a.url + "/" + id
	return a
}

type idLabel struct {
	url string
}

func createIdLabel(m model) idLabel {
	return idLabel{
		url: m.getURL() + "/idLabels",
	}
}

func (i idLabel) ID(id string) idLabel {
	i.url = i.url + "/" + id
	return i
}

type idMember struct {
	url string
}

func createIdMember(m model) idMember {
	return idMember{
		url: m.getURL() + "/idMembers",
	}
}

func (i idMember) ID(id string) idMember {
	i.url = i.url + "/" + id
	return i
}

type membersVoted struct {
	url string
}

func createMembersVoted(m model) membersVoted {
	return membersVoted{
		url: m.getURL() + "/membersVoted",
	}
}

func (i membersVoted) ID(id string) membersVoted {
	i.url = i.url + "/" + id
	return i
}

type sticker struct {
	url string
}

func createSticker(m model) sticker {
	return sticker{
		url: m.getURL() + "/stickers",
	}
}

func (i sticker) ID(id string) sticker {
	i.url = i.url + "/" + id
	return i
}

type filterCard struct {
	url string
	All,
	Closed,
	Open,
	None,
	Visible staticField
}

func createFilterCard(m model) filterCard {
	cURL := m.getURL() + "/cards"
	return filterCard{
		url:     cURL,
		All:     staticField(cURL + "/all"),
		Closed:  staticField(cURL + "/closed"),
		Open:    staticField(cURL + "/open"),
		None:    staticField(cURL + "/none"),
		Visible: staticField(cURL + "/visible"),
	}
}

func (f filterCard) ID(id string) filterCard {
	fURL := f.url + "/" + id
	f.url = fURL
	f.All = staticField(fURL + "/all")
	f.Closed = staticField(fURL + "/closed")
	f.None = staticField(fURL + "/none")
	f.Open = staticField(fURL + "/open")
	f.Visible = staticField(fURL + "/visible")
	return f
}

func (c filterCard) getURL() string {
	return c.url
}

type baseCard struct {
	filterCard
	url string
	//static fields
	Badges,
	CheckItemStates,
	Closed,
	DateLastActivity,
	Desc,
	DescData,
	Due,
	Email,
	IdAttachmentCover,
	IdBoard,
	IdChecklists,
	IdList,
	IdMembersVoted,
	IdShort,
	ManualCoverAttachment,
	Members,
	Name,
	Pos,
	ShortLink,
	ShortUrl,
	Subscribed,
	Url staticField
}

func createBaseCard(m model) baseCard {
	var cardURL string

	//Handle special case for type
	switch m.(type) {
	case action:
		cardURL = m.getURL() + "/card"
		break
	default:
		cardURL = m.getURL() + "/cards"
	}

	c := baseCard{}
	c.filterCard = createFilterCard(m)
	c.url = cardURL
	c.Badges = staticField(cardURL + "/badges")
	c.CheckItemStates = staticField(cardURL + "/checkItemStates")
	c.Closed = staticField(cardURL + "/closed")
	c.DateLastActivity = staticField(cardURL + "/dateLastActivity")
	c.Desc = staticField(cardURL + "/desc")
	c.DescData = staticField(cardURL + "/descData")
	c.Due = staticField(cardURL + "/due")
	c.Email = staticField(cardURL + "/email")
	c.IdAttachmentCover = staticField(cardURL + "/idAttachmentCover")
	c.IdBoard = staticField(cardURL + "/idBoard")
	c.IdChecklists = staticField(cardURL + "/idChecklist")
	c.IdList = staticField(cardURL + "/idList")
	c.IdMembersVoted = staticField(cardURL + "/idMembersVoted")
	c.IdShort = staticField(cardURL + "/idShort")
	c.ManualCoverAttachment = staticField(cardURL + "/manualCoverAttachment")
	c.Members = staticField(cardURL + "/members")
	c.Name = staticField(cardURL + "/name")
	c.Pos = staticField(cardURL + "/pos")
	c.ShortLink = staticField(cardURL + "/shortLink")
	c.ShortUrl = staticField(cardURL + "/shortUrl")
	c.Subscribed = staticField(cardURL + "/subscribed")
	c.Url = staticField(cardURL + "/url")
	return c
}

func (c baseCard) getURL() string {
	return c.url
}

type card struct {
	baseCard
	url string
	//submodels
	Actions      baseAction
	Attachments  attachment
	Board        baseBoard
	Checklist    checklist
	Checklists   checklists
	IdLabels     idLabel
	IdMembers    idMember
	Labels       label
	List         list
	MembersVoted membersVoted
	Stickers     sticker
}

func createCard(m model) card {
	return card{
		url: m.getURL() + "/cards",
	}
}

func (c card) ID(id string) card {
	cardURL := c.url + "/" + id
	c.url = cardURL
	c.Badges = staticField(cardURL + "/badges")
	c.CheckItemStates = staticField(cardURL + "/checkItemStates")
	c.Closed = staticField(cardURL + "/closed")
	c.DateLastActivity = staticField(cardURL + "/dateLastActivity")
	c.Desc = staticField(cardURL + "/desc")
	c.DescData = staticField(cardURL + "/descData")
	c.Due = staticField(cardURL + "/due")
	c.Email = staticField(cardURL + "/email")
	c.IdAttachmentCover = staticField(cardURL + "/idAttachmentCover")
	c.IdBoard = staticField(cardURL + "/idBoard")
	c.IdChecklists = staticField(cardURL + "/idChecklist")
	c.IdList = staticField(cardURL + "/idList")
	c.IdMembersVoted = staticField(cardURL + "/idMembersVoted")
	c.IdShort = staticField(cardURL + "/idShort")
	c.ManualCoverAttachment = staticField(cardURL + "/manualCoverAttachment")
	c.Members = staticField(cardURL + "/members")
	c.Name = staticField(cardURL + "/name")
	c.Pos = staticField(cardURL + "/pos")
	c.ShortLink = staticField(cardURL + "/shortLink")
	c.ShortUrl = staticField(cardURL + "/shortUrl")
	c.Subscribed = staticField(cardURL + "/subscribed")
	c.Url = staticField(cardURL + "/url")
	c.Actions = createBaseAction(c)
	c.Attachments = createAttachment(c)
	c.Board = createBaseBoard(c)
	c.Checklists = createChecklists(c)
	c.IdLabels = createIdLabel(c)
	c.IdMembers = createIdMember(c)
	c.Labels = createLabel(c)
	return c
}

func (c card) getURL() string {
	return c.url
}

var Card = card{
	url: "/cards",
}
