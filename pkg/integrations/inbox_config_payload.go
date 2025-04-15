package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/util"
)

type InboxConfigPayload struct {
	Images *util.Nullable[Images] `json:"images,omitempty" required:"true"`
	Locale *util.Nullable[string] `json:"locale,omitempty" required:"true" minLength:"2"`
	Theme  *util.Nullable[Theme]  `json:"theme,omitempty" required:"true"`
}

func (i *InboxConfigPayload) GetImages() *util.Nullable[Images] {
	if i == nil {
		return nil
	}
	return i.Images
}

func (i *InboxConfigPayload) SetImages(images util.Nullable[Images]) {
	i.Images = &images
}

func (i *InboxConfigPayload) SetImagesNull() {
	i.Images = &util.Nullable[Images]{IsNull: true}
}

func (i *InboxConfigPayload) GetLocale() *util.Nullable[string] {
	if i == nil {
		return nil
	}
	return i.Locale
}

func (i *InboxConfigPayload) SetLocale(locale util.Nullable[string]) {
	i.Locale = &locale
}

func (i *InboxConfigPayload) SetLocaleNull() {
	i.Locale = &util.Nullable[string]{IsNull: true}
}

func (i *InboxConfigPayload) GetTheme() *util.Nullable[Theme] {
	if i == nil {
		return nil
	}
	return i.Theme
}

func (i *InboxConfigPayload) SetTheme(theme util.Nullable[Theme]) {
	i.Theme = &theme
}

func (i *InboxConfigPayload) SetThemeNull() {
	i.Theme = &util.Nullable[Theme]{IsNull: true}
}

func (i InboxConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InboxConfigPayload to string"
	}
	return string(jsonData)
}

type Images struct {
	EmptyInboxUrl *string `json:"emptyInboxUrl,omitempty" required:"true"`
}

func (i *Images) GetEmptyInboxUrl() *string {
	if i == nil {
		return nil
	}
	return i.EmptyInboxUrl
}

func (i *Images) SetEmptyInboxUrl(emptyInboxUrl string) {
	i.EmptyInboxUrl = &emptyInboxUrl
}

func (i Images) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: Images to string"
	}
	return string(jsonData)
}

type Theme struct {
	Banner       *Banner       `json:"banner,omitempty"`
	Dialog       *Dialog       `json:"dialog,omitempty"`
	Footer       *Footer       `json:"footer,omitempty"`
	Header       *Header       `json:"header,omitempty"`
	Icon         *Icon         `json:"icon,omitempty"`
	Notification *Notification `json:"notification,omitempty"`
	UnseenBadge  *UnseenBadge  `json:"unseenBadge,omitempty"`
}

func (t *Theme) GetBanner() *Banner {
	if t == nil {
		return nil
	}
	return t.Banner
}

func (t *Theme) SetBanner(banner Banner) {
	t.Banner = &banner
}

func (t *Theme) GetDialog() *Dialog {
	if t == nil {
		return nil
	}
	return t.Dialog
}

func (t *Theme) SetDialog(dialog Dialog) {
	t.Dialog = &dialog
}

func (t *Theme) GetFooter() *Footer {
	if t == nil {
		return nil
	}
	return t.Footer
}

func (t *Theme) SetFooter(footer Footer) {
	t.Footer = &footer
}

func (t *Theme) GetHeader() *Header {
	if t == nil {
		return nil
	}
	return t.Header
}

func (t *Theme) SetHeader(header Header) {
	t.Header = &header
}

func (t *Theme) GetIcon() *Icon {
	if t == nil {
		return nil
	}
	return t.Icon
}

func (t *Theme) SetIcon(icon Icon) {
	t.Icon = &icon
}

func (t *Theme) GetNotification() *Notification {
	if t == nil {
		return nil
	}
	return t.Notification
}

func (t *Theme) SetNotification(notification Notification) {
	t.Notification = &notification
}

func (t *Theme) GetUnseenBadge() *UnseenBadge {
	if t == nil {
		return nil
	}
	return t.UnseenBadge
}

func (t *Theme) SetUnseenBadge(unseenBadge UnseenBadge) {
	t.UnseenBadge = &unseenBadge
}

func (t Theme) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: Theme to string"
	}
	return string(jsonData)
}

type Banner struct {
	BackgroundColor   *string  `json:"backgroundColor,omitempty" required:"true"`
	BackgroundOpacity *float64 `json:"backgroundOpacity,omitempty"`
	FontSize          *string  `json:"fontSize,omitempty" required:"true"`
	TextColor         *string  `json:"textColor,omitempty" required:"true"`
}

func (b *Banner) GetBackgroundColor() *string {
	if b == nil {
		return nil
	}
	return b.BackgroundColor
}

func (b *Banner) SetBackgroundColor(backgroundColor string) {
	b.BackgroundColor = &backgroundColor
}

func (b *Banner) GetBackgroundOpacity() *float64 {
	if b == nil {
		return nil
	}
	return b.BackgroundOpacity
}

func (b *Banner) SetBackgroundOpacity(backgroundOpacity float64) {
	b.BackgroundOpacity = &backgroundOpacity
}

func (b *Banner) GetFontSize() *string {
	if b == nil {
		return nil
	}
	return b.FontSize
}

func (b *Banner) SetFontSize(fontSize string) {
	b.FontSize = &fontSize
}

func (b *Banner) GetTextColor() *string {
	if b == nil {
		return nil
	}
	return b.TextColor
}

func (b *Banner) SetTextColor(textColor string) {
	b.TextColor = &textColor
}

func (b Banner) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: Banner to string"
	}
	return string(jsonData)
}

type Dialog struct {
	AccentColor     *string `json:"accentColor,omitempty" required:"true"`
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	TextColor       *string `json:"textColor,omitempty" required:"true"`
}

func (d *Dialog) GetAccentColor() *string {
	if d == nil {
		return nil
	}
	return d.AccentColor
}

func (d *Dialog) SetAccentColor(accentColor string) {
	d.AccentColor = &accentColor
}

func (d *Dialog) GetBackgroundColor() *string {
	if d == nil {
		return nil
	}
	return d.BackgroundColor
}

func (d *Dialog) SetBackgroundColor(backgroundColor string) {
	d.BackgroundColor = &backgroundColor
}

func (d *Dialog) GetTextColor() *string {
	if d == nil {
		return nil
	}
	return d.TextColor
}

func (d *Dialog) SetTextColor(textColor string) {
	d.TextColor = &textColor
}

func (d Dialog) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: Dialog to string"
	}
	return string(jsonData)
}

type Footer struct {
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	BorderRadius    *string `json:"borderRadius,omitempty" required:"true"`
	FontSize        *string `json:"fontSize,omitempty" required:"true"`
	TextColor       *string `json:"textColor,omitempty" required:"true"`
}

func (f *Footer) GetBackgroundColor() *string {
	if f == nil {
		return nil
	}
	return f.BackgroundColor
}

func (f *Footer) SetBackgroundColor(backgroundColor string) {
	f.BackgroundColor = &backgroundColor
}

func (f *Footer) GetBorderRadius() *string {
	if f == nil {
		return nil
	}
	return f.BorderRadius
}

func (f *Footer) SetBorderRadius(borderRadius string) {
	f.BorderRadius = &borderRadius
}

func (f *Footer) GetFontSize() *string {
	if f == nil {
		return nil
	}
	return f.FontSize
}

func (f *Footer) SetFontSize(fontSize string) {
	f.FontSize = &fontSize
}

func (f *Footer) GetTextColor() *string {
	if f == nil {
		return nil
	}
	return f.TextColor
}

func (f *Footer) SetTextColor(textColor string) {
	f.TextColor = &textColor
}

func (f Footer) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: Footer to string"
	}
	return string(jsonData)
}

type Header struct {
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	BorderRadius    *string `json:"borderRadius,omitempty" required:"true"`
	FontFamily      *string `json:"fontFamily,omitempty" required:"true"`
	FontSize        *string `json:"fontSize,omitempty" required:"true"`
	TextColor       *string `json:"textColor,omitempty" required:"true"`
}

func (h *Header) GetBackgroundColor() *string {
	if h == nil {
		return nil
	}
	return h.BackgroundColor
}

func (h *Header) SetBackgroundColor(backgroundColor string) {
	h.BackgroundColor = &backgroundColor
}

func (h *Header) GetBorderRadius() *string {
	if h == nil {
		return nil
	}
	return h.BorderRadius
}

func (h *Header) SetBorderRadius(borderRadius string) {
	h.BorderRadius = &borderRadius
}

func (h *Header) GetFontFamily() *string {
	if h == nil {
		return nil
	}
	return h.FontFamily
}

func (h *Header) SetFontFamily(fontFamily string) {
	h.FontFamily = &fontFamily
}

func (h *Header) GetFontSize() *string {
	if h == nil {
		return nil
	}
	return h.FontSize
}

func (h *Header) SetFontSize(fontSize string) {
	h.FontSize = &fontSize
}

func (h *Header) GetTextColor() *string {
	if h == nil {
		return nil
	}
	return h.TextColor
}

func (h *Header) SetTextColor(textColor string) {
	h.TextColor = &textColor
}

func (h Header) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: Header to string"
	}
	return string(jsonData)
}

type Icon struct {
	BorderColor *string `json:"borderColor,omitempty" required:"true"`
	Width       *string `json:"width,omitempty" required:"true"`
}

func (i *Icon) GetBorderColor() *string {
	if i == nil {
		return nil
	}
	return i.BorderColor
}

func (i *Icon) SetBorderColor(borderColor string) {
	i.BorderColor = &borderColor
}

func (i *Icon) GetWidth() *string {
	if i == nil {
		return nil
	}
	return i.Width
}

func (i *Icon) SetWidth(width string) {
	i.Width = &width
}

func (i Icon) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: Icon to string"
	}
	return string(jsonData)
}

type Notification struct {
	Default_ *Default_ `json:"default,omitempty" required:"true"`
	Unread   *Unread   `json:"unread,omitempty" required:"true"`
	Unseen   *Unseen   `json:"unseen,omitempty" required:"true"`
}

func (n *Notification) GetDefault_() *Default_ {
	if n == nil {
		return nil
	}
	return n.Default_
}

func (n *Notification) SetDefault_(default_ Default_) {
	n.Default_ = &default_
}

func (n *Notification) GetUnread() *Unread {
	if n == nil {
		return nil
	}
	return n.Unread
}

func (n *Notification) SetUnread(unread Unread) {
	n.Unread = &unread
}

func (n *Notification) GetUnseen() *Unseen {
	if n == nil {
		return nil
	}
	return n.Unseen
}

func (n *Notification) SetUnseen(unseen Unseen) {
	n.Unseen = &unseen
}

func (n Notification) String() string {
	jsonData, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return "error converting struct: Notification to string"
	}
	return string(jsonData)
}

type Default_ struct {
	BackgroundColor *string       `json:"backgroundColor,omitempty" required:"true"`
	BorderRadius    *string       `json:"borderRadius,omitempty" required:"true"`
	FontFamily      *string       `json:"fontFamily,omitempty" required:"true"`
	FontSize        *string       `json:"fontSize,omitempty" required:"true"`
	Hover           *DefaultHover `json:"hover,omitempty"`
	Margin          *string       `json:"margin,omitempty" required:"true"`
	State           *DefaultState `json:"state,omitempty"`
	TextColor       *string       `json:"textColor,omitempty" required:"true"`
}

func (d *Default_) GetBackgroundColor() *string {
	if d == nil {
		return nil
	}
	return d.BackgroundColor
}

func (d *Default_) SetBackgroundColor(backgroundColor string) {
	d.BackgroundColor = &backgroundColor
}

func (d *Default_) GetBorderRadius() *string {
	if d == nil {
		return nil
	}
	return d.BorderRadius
}

func (d *Default_) SetBorderRadius(borderRadius string) {
	d.BorderRadius = &borderRadius
}

func (d *Default_) GetFontFamily() *string {
	if d == nil {
		return nil
	}
	return d.FontFamily
}

func (d *Default_) SetFontFamily(fontFamily string) {
	d.FontFamily = &fontFamily
}

func (d *Default_) GetFontSize() *string {
	if d == nil {
		return nil
	}
	return d.FontSize
}

func (d *Default_) SetFontSize(fontSize string) {
	d.FontSize = &fontSize
}

func (d *Default_) GetHover() *DefaultHover {
	if d == nil {
		return nil
	}
	return d.Hover
}

func (d *Default_) SetHover(hover DefaultHover) {
	d.Hover = &hover
}

func (d *Default_) GetMargin() *string {
	if d == nil {
		return nil
	}
	return d.Margin
}

func (d *Default_) SetMargin(margin string) {
	d.Margin = &margin
}

func (d *Default_) GetState() *DefaultState {
	if d == nil {
		return nil
	}
	return d.State
}

func (d *Default_) SetState(state DefaultState) {
	d.State = &state
}

func (d *Default_) GetTextColor() *string {
	if d == nil {
		return nil
	}
	return d.TextColor
}

func (d *Default_) SetTextColor(textColor string) {
	d.TextColor = &textColor
}

func (d Default_) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: Default_ to string"
	}
	return string(jsonData)
}

type DefaultHover struct {
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
}

func (d *DefaultHover) GetBackgroundColor() *string {
	if d == nil {
		return nil
	}
	return d.BackgroundColor
}

func (d *DefaultHover) SetBackgroundColor(backgroundColor string) {
	d.BackgroundColor = &backgroundColor
}

func (d DefaultHover) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DefaultHover to string"
	}
	return string(jsonData)
}

type DefaultState struct {
	Color *string `json:"color,omitempty" required:"true"`
}

func (d *DefaultState) GetColor() *string {
	if d == nil {
		return nil
	}
	return d.Color
}

func (d *DefaultState) SetColor(color string) {
	d.Color = &color
}

func (d DefaultState) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DefaultState to string"
	}
	return string(jsonData)
}

type Unread struct {
	BackgroundColor *string      `json:"backgroundColor,omitempty" required:"true"`
	Hover           *UnreadHover `json:"hover,omitempty"`
	State           *UnreadState `json:"state,omitempty"`
	TextColor       *string      `json:"textColor,omitempty" required:"true"`
}

func (u *Unread) GetBackgroundColor() *string {
	if u == nil {
		return nil
	}
	return u.BackgroundColor
}

func (u *Unread) SetBackgroundColor(backgroundColor string) {
	u.BackgroundColor = &backgroundColor
}

func (u *Unread) GetHover() *UnreadHover {
	if u == nil {
		return nil
	}
	return u.Hover
}

func (u *Unread) SetHover(hover UnreadHover) {
	u.Hover = &hover
}

func (u *Unread) GetState() *UnreadState {
	if u == nil {
		return nil
	}
	return u.State
}

func (u *Unread) SetState(state UnreadState) {
	u.State = &state
}

func (u *Unread) GetTextColor() *string {
	if u == nil {
		return nil
	}
	return u.TextColor
}

func (u *Unread) SetTextColor(textColor string) {
	u.TextColor = &textColor
}

func (u Unread) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: Unread to string"
	}
	return string(jsonData)
}

type UnreadHover struct {
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
}

func (u *UnreadHover) GetBackgroundColor() *string {
	if u == nil {
		return nil
	}
	return u.BackgroundColor
}

func (u *UnreadHover) SetBackgroundColor(backgroundColor string) {
	u.BackgroundColor = &backgroundColor
}

func (u UnreadHover) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UnreadHover to string"
	}
	return string(jsonData)
}

type UnreadState struct {
	Color *string `json:"color,omitempty" required:"true"`
}

func (u *UnreadState) GetColor() *string {
	if u == nil {
		return nil
	}
	return u.Color
}

func (u *UnreadState) SetColor(color string) {
	u.Color = &color
}

func (u UnreadState) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UnreadState to string"
	}
	return string(jsonData)
}

type Unseen struct {
	BackgroundColor *string      `json:"backgroundColor,omitempty" required:"true"`
	Hover           *UnseenHover `json:"hover,omitempty"`
	State           *UnseenState `json:"state,omitempty"`
	TextColor       *string      `json:"textColor,omitempty" required:"true"`
}

func (u *Unseen) GetBackgroundColor() *string {
	if u == nil {
		return nil
	}
	return u.BackgroundColor
}

func (u *Unseen) SetBackgroundColor(backgroundColor string) {
	u.BackgroundColor = &backgroundColor
}

func (u *Unseen) GetHover() *UnseenHover {
	if u == nil {
		return nil
	}
	return u.Hover
}

func (u *Unseen) SetHover(hover UnseenHover) {
	u.Hover = &hover
}

func (u *Unseen) GetState() *UnseenState {
	if u == nil {
		return nil
	}
	return u.State
}

func (u *Unseen) SetState(state UnseenState) {
	u.State = &state
}

func (u *Unseen) GetTextColor() *string {
	if u == nil {
		return nil
	}
	return u.TextColor
}

func (u *Unseen) SetTextColor(textColor string) {
	u.TextColor = &textColor
}

func (u Unseen) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: Unseen to string"
	}
	return string(jsonData)
}

type UnseenHover struct {
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
}

func (u *UnseenHover) GetBackgroundColor() *string {
	if u == nil {
		return nil
	}
	return u.BackgroundColor
}

func (u *UnseenHover) SetBackgroundColor(backgroundColor string) {
	u.BackgroundColor = &backgroundColor
}

func (u UnseenHover) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UnseenHover to string"
	}
	return string(jsonData)
}

type UnseenState struct {
	Color *string `json:"color,omitempty" required:"true"`
}

func (u *UnseenState) GetColor() *string {
	if u == nil {
		return nil
	}
	return u.Color
}

func (u *UnseenState) SetColor(color string) {
	u.Color = &color
}

func (u UnseenState) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UnseenState to string"
	}
	return string(jsonData)
}

type UnseenBadge struct {
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
}

func (u *UnseenBadge) GetBackgroundColor() *string {
	if u == nil {
		return nil
	}
	return u.BackgroundColor
}

func (u *UnseenBadge) SetBackgroundColor(backgroundColor string) {
	u.BackgroundColor = &backgroundColor
}

func (u UnseenBadge) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UnseenBadge to string"
	}
	return string(jsonData)
}
