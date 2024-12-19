package integrations

import (
	"encoding/json"
)

type InboxConfig struct {
	Images  *Images `json:"images,omitempty" required:"true"`
	Locale  *string `json:"locale,omitempty" required:"true" minLength:"2"`
	Theme   *Theme  `json:"theme,omitempty" required:"true"`
	touched map[string]bool
}

func (i *InboxConfig) GetImages() *Images {
	if i == nil {
		return nil
	}
	return i.Images
}

func (i *InboxConfig) SetImages(images Images) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Images"] = true
	i.Images = &images
}

func (i *InboxConfig) SetImagesNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Images"] = true
	i.Images = nil
}

func (i *InboxConfig) GetLocale() *string {
	if i == nil {
		return nil
	}
	return i.Locale
}

func (i *InboxConfig) SetLocale(locale string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Locale"] = true
	i.Locale = &locale
}

func (i *InboxConfig) SetLocaleNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Locale"] = true
	i.Locale = nil
}

func (i *InboxConfig) GetTheme() *Theme {
	if i == nil {
		return nil
	}
	return i.Theme
}

func (i *InboxConfig) SetTheme(theme Theme) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Theme"] = true
	i.Theme = &theme
}

func (i *InboxConfig) SetThemeNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Theme"] = true
	i.Theme = nil
}
func (i InboxConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if i.touched["Images"] && i.Images == nil {
		data["images"] = nil
	} else if i.Images != nil {
		data["images"] = i.Images
	}

	if i.touched["Locale"] && i.Locale == nil {
		data["locale"] = nil
	} else if i.Locale != nil {
		data["locale"] = i.Locale
	}

	if i.touched["Theme"] && i.Theme == nil {
		data["theme"] = nil
	} else if i.Theme != nil {
		data["theme"] = i.Theme
	}

	return json.Marshal(data)
}

type Images struct {
	EmptyInboxUrl *string `json:"emptyInboxUrl,omitempty" required:"true"`
	touched       map[string]bool
}

func (i *Images) GetEmptyInboxUrl() *string {
	if i == nil {
		return nil
	}
	return i.EmptyInboxUrl
}

func (i *Images) SetEmptyInboxUrl(emptyInboxUrl string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["EmptyInboxUrl"] = true
	i.EmptyInboxUrl = &emptyInboxUrl
}

func (i *Images) SetEmptyInboxUrlNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["EmptyInboxUrl"] = true
	i.EmptyInboxUrl = nil
}
func (i Images) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if i.touched["EmptyInboxUrl"] && i.EmptyInboxUrl == nil {
		data["emptyInboxUrl"] = nil
	} else if i.EmptyInboxUrl != nil {
		data["emptyInboxUrl"] = i.EmptyInboxUrl
	}

	return json.Marshal(data)
}

type Theme struct {
	Banner       *Banner       `json:"banner,omitempty"`
	Dialog       *Dialog       `json:"dialog,omitempty"`
	Footer       *Footer       `json:"footer,omitempty"`
	Header       *Header       `json:"header,omitempty"`
	Icon         *Icon         `json:"icon,omitempty"`
	Notification *Notification `json:"notification,omitempty"`
	UnseenBadge  *UnseenBadge  `json:"unseenBadge,omitempty"`
	touched      map[string]bool
}

func (t *Theme) GetBanner() *Banner {
	if t == nil {
		return nil
	}
	return t.Banner
}

func (t *Theme) SetBanner(banner Banner) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Banner"] = true
	t.Banner = &banner
}

func (t *Theme) SetBannerNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Banner"] = true
	t.Banner = nil
}

func (t *Theme) GetDialog() *Dialog {
	if t == nil {
		return nil
	}
	return t.Dialog
}

func (t *Theme) SetDialog(dialog Dialog) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Dialog"] = true
	t.Dialog = &dialog
}

func (t *Theme) SetDialogNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Dialog"] = true
	t.Dialog = nil
}

func (t *Theme) GetFooter() *Footer {
	if t == nil {
		return nil
	}
	return t.Footer
}

func (t *Theme) SetFooter(footer Footer) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Footer"] = true
	t.Footer = &footer
}

func (t *Theme) SetFooterNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Footer"] = true
	t.Footer = nil
}

func (t *Theme) GetHeader() *Header {
	if t == nil {
		return nil
	}
	return t.Header
}

func (t *Theme) SetHeader(header Header) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Header"] = true
	t.Header = &header
}

func (t *Theme) SetHeaderNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Header"] = true
	t.Header = nil
}

func (t *Theme) GetIcon() *Icon {
	if t == nil {
		return nil
	}
	return t.Icon
}

func (t *Theme) SetIcon(icon Icon) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Icon"] = true
	t.Icon = &icon
}

func (t *Theme) SetIconNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Icon"] = true
	t.Icon = nil
}

func (t *Theme) GetNotification() *Notification {
	if t == nil {
		return nil
	}
	return t.Notification
}

func (t *Theme) SetNotification(notification Notification) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Notification"] = true
	t.Notification = &notification
}

func (t *Theme) SetNotificationNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Notification"] = true
	t.Notification = nil
}

func (t *Theme) GetUnseenBadge() *UnseenBadge {
	if t == nil {
		return nil
	}
	return t.UnseenBadge
}

func (t *Theme) SetUnseenBadge(unseenBadge UnseenBadge) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["UnseenBadge"] = true
	t.UnseenBadge = &unseenBadge
}

func (t *Theme) SetUnseenBadgeNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["UnseenBadge"] = true
	t.UnseenBadge = nil
}
func (t Theme) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["Banner"] && t.Banner == nil {
		data["banner"] = nil
	} else if t.Banner != nil {
		data["banner"] = t.Banner
	}

	if t.touched["Dialog"] && t.Dialog == nil {
		data["dialog"] = nil
	} else if t.Dialog != nil {
		data["dialog"] = t.Dialog
	}

	if t.touched["Footer"] && t.Footer == nil {
		data["footer"] = nil
	} else if t.Footer != nil {
		data["footer"] = t.Footer
	}

	if t.touched["Header"] && t.Header == nil {
		data["header"] = nil
	} else if t.Header != nil {
		data["header"] = t.Header
	}

	if t.touched["Icon"] && t.Icon == nil {
		data["icon"] = nil
	} else if t.Icon != nil {
		data["icon"] = t.Icon
	}

	if t.touched["Notification"] && t.Notification == nil {
		data["notification"] = nil
	} else if t.Notification != nil {
		data["notification"] = t.Notification
	}

	if t.touched["UnseenBadge"] && t.UnseenBadge == nil {
		data["unseenBadge"] = nil
	} else if t.UnseenBadge != nil {
		data["unseenBadge"] = t.UnseenBadge
	}

	return json.Marshal(data)
}

type Banner struct {
	BackgroundColor   *string  `json:"backgroundColor,omitempty" required:"true"`
	BackgroundOpacity *float64 `json:"backgroundOpacity,omitempty"`
	FontSize          *string  `json:"fontSize,omitempty" required:"true"`
	TextColor         *string  `json:"textColor,omitempty" required:"true"`
	touched           map[string]bool
}

func (b *Banner) GetBackgroundColor() *string {
	if b == nil {
		return nil
	}
	return b.BackgroundColor
}

func (b *Banner) SetBackgroundColor(backgroundColor string) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["BackgroundColor"] = true
	b.BackgroundColor = &backgroundColor
}

func (b *Banner) SetBackgroundColorNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["BackgroundColor"] = true
	b.BackgroundColor = nil
}

func (b *Banner) GetBackgroundOpacity() *float64 {
	if b == nil {
		return nil
	}
	return b.BackgroundOpacity
}

func (b *Banner) SetBackgroundOpacity(backgroundOpacity float64) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["BackgroundOpacity"] = true
	b.BackgroundOpacity = &backgroundOpacity
}

func (b *Banner) SetBackgroundOpacityNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["BackgroundOpacity"] = true
	b.BackgroundOpacity = nil
}

func (b *Banner) GetFontSize() *string {
	if b == nil {
		return nil
	}
	return b.FontSize
}

func (b *Banner) SetFontSize(fontSize string) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["FontSize"] = true
	b.FontSize = &fontSize
}

func (b *Banner) SetFontSizeNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["FontSize"] = true
	b.FontSize = nil
}

func (b *Banner) GetTextColor() *string {
	if b == nil {
		return nil
	}
	return b.TextColor
}

func (b *Banner) SetTextColor(textColor string) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["TextColor"] = true
	b.TextColor = &textColor
}

func (b *Banner) SetTextColorNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["TextColor"] = true
	b.TextColor = nil
}
func (b Banner) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if b.touched["BackgroundColor"] && b.BackgroundColor == nil {
		data["backgroundColor"] = nil
	} else if b.BackgroundColor != nil {
		data["backgroundColor"] = b.BackgroundColor
	}

	if b.touched["BackgroundOpacity"] && b.BackgroundOpacity == nil {
		data["backgroundOpacity"] = nil
	} else if b.BackgroundOpacity != nil {
		data["backgroundOpacity"] = b.BackgroundOpacity
	}

	if b.touched["FontSize"] && b.FontSize == nil {
		data["fontSize"] = nil
	} else if b.FontSize != nil {
		data["fontSize"] = b.FontSize
	}

	if b.touched["TextColor"] && b.TextColor == nil {
		data["textColor"] = nil
	} else if b.TextColor != nil {
		data["textColor"] = b.TextColor
	}

	return json.Marshal(data)
}

type Dialog struct {
	AccentColor     *string `json:"accentColor,omitempty" required:"true"`
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	TextColor       *string `json:"textColor,omitempty" required:"true"`
	touched         map[string]bool
}

func (d *Dialog) GetAccentColor() *string {
	if d == nil {
		return nil
	}
	return d.AccentColor
}

func (d *Dialog) SetAccentColor(accentColor string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["AccentColor"] = true
	d.AccentColor = &accentColor
}

func (d *Dialog) SetAccentColorNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["AccentColor"] = true
	d.AccentColor = nil
}

func (d *Dialog) GetBackgroundColor() *string {
	if d == nil {
		return nil
	}
	return d.BackgroundColor
}

func (d *Dialog) SetBackgroundColor(backgroundColor string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["BackgroundColor"] = true
	d.BackgroundColor = &backgroundColor
}

func (d *Dialog) SetBackgroundColorNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["BackgroundColor"] = true
	d.BackgroundColor = nil
}

func (d *Dialog) GetTextColor() *string {
	if d == nil {
		return nil
	}
	return d.TextColor
}

func (d *Dialog) SetTextColor(textColor string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["TextColor"] = true
	d.TextColor = &textColor
}

func (d *Dialog) SetTextColorNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["TextColor"] = true
	d.TextColor = nil
}
func (d Dialog) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["AccentColor"] && d.AccentColor == nil {
		data["accentColor"] = nil
	} else if d.AccentColor != nil {
		data["accentColor"] = d.AccentColor
	}

	if d.touched["BackgroundColor"] && d.BackgroundColor == nil {
		data["backgroundColor"] = nil
	} else if d.BackgroundColor != nil {
		data["backgroundColor"] = d.BackgroundColor
	}

	if d.touched["TextColor"] && d.TextColor == nil {
		data["textColor"] = nil
	} else if d.TextColor != nil {
		data["textColor"] = d.TextColor
	}

	return json.Marshal(data)
}

type Footer struct {
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	BorderRadius    *string `json:"borderRadius,omitempty" required:"true"`
	FontSize        *string `json:"fontSize,omitempty" required:"true"`
	TextColor       *string `json:"textColor,omitempty" required:"true"`
	touched         map[string]bool
}

func (f *Footer) GetBackgroundColor() *string {
	if f == nil {
		return nil
	}
	return f.BackgroundColor
}

func (f *Footer) SetBackgroundColor(backgroundColor string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["BackgroundColor"] = true
	f.BackgroundColor = &backgroundColor
}

func (f *Footer) SetBackgroundColorNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["BackgroundColor"] = true
	f.BackgroundColor = nil
}

func (f *Footer) GetBorderRadius() *string {
	if f == nil {
		return nil
	}
	return f.BorderRadius
}

func (f *Footer) SetBorderRadius(borderRadius string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["BorderRadius"] = true
	f.BorderRadius = &borderRadius
}

func (f *Footer) SetBorderRadiusNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["BorderRadius"] = true
	f.BorderRadius = nil
}

func (f *Footer) GetFontSize() *string {
	if f == nil {
		return nil
	}
	return f.FontSize
}

func (f *Footer) SetFontSize(fontSize string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["FontSize"] = true
	f.FontSize = &fontSize
}

func (f *Footer) SetFontSizeNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["FontSize"] = true
	f.FontSize = nil
}

func (f *Footer) GetTextColor() *string {
	if f == nil {
		return nil
	}
	return f.TextColor
}

func (f *Footer) SetTextColor(textColor string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["TextColor"] = true
	f.TextColor = &textColor
}

func (f *Footer) SetTextColorNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["TextColor"] = true
	f.TextColor = nil
}
func (f Footer) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if f.touched["BackgroundColor"] && f.BackgroundColor == nil {
		data["backgroundColor"] = nil
	} else if f.BackgroundColor != nil {
		data["backgroundColor"] = f.BackgroundColor
	}

	if f.touched["BorderRadius"] && f.BorderRadius == nil {
		data["borderRadius"] = nil
	} else if f.BorderRadius != nil {
		data["borderRadius"] = f.BorderRadius
	}

	if f.touched["FontSize"] && f.FontSize == nil {
		data["fontSize"] = nil
	} else if f.FontSize != nil {
		data["fontSize"] = f.FontSize
	}

	if f.touched["TextColor"] && f.TextColor == nil {
		data["textColor"] = nil
	} else if f.TextColor != nil {
		data["textColor"] = f.TextColor
	}

	return json.Marshal(data)
}

type Header struct {
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	BorderRadius    *string `json:"borderRadius,omitempty" required:"true"`
	FontFamily      *string `json:"fontFamily,omitempty" required:"true"`
	FontSize        *string `json:"fontSize,omitempty" required:"true"`
	TextColor       *string `json:"textColor,omitempty" required:"true"`
	touched         map[string]bool
}

func (h *Header) GetBackgroundColor() *string {
	if h == nil {
		return nil
	}
	return h.BackgroundColor
}

func (h *Header) SetBackgroundColor(backgroundColor string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["BackgroundColor"] = true
	h.BackgroundColor = &backgroundColor
}

func (h *Header) SetBackgroundColorNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["BackgroundColor"] = true
	h.BackgroundColor = nil
}

func (h *Header) GetBorderRadius() *string {
	if h == nil {
		return nil
	}
	return h.BorderRadius
}

func (h *Header) SetBorderRadius(borderRadius string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["BorderRadius"] = true
	h.BorderRadius = &borderRadius
}

func (h *Header) SetBorderRadiusNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["BorderRadius"] = true
	h.BorderRadius = nil
}

func (h *Header) GetFontFamily() *string {
	if h == nil {
		return nil
	}
	return h.FontFamily
}

func (h *Header) SetFontFamily(fontFamily string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["FontFamily"] = true
	h.FontFamily = &fontFamily
}

func (h *Header) SetFontFamilyNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["FontFamily"] = true
	h.FontFamily = nil
}

func (h *Header) GetFontSize() *string {
	if h == nil {
		return nil
	}
	return h.FontSize
}

func (h *Header) SetFontSize(fontSize string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["FontSize"] = true
	h.FontSize = &fontSize
}

func (h *Header) SetFontSizeNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["FontSize"] = true
	h.FontSize = nil
}

func (h *Header) GetTextColor() *string {
	if h == nil {
		return nil
	}
	return h.TextColor
}

func (h *Header) SetTextColor(textColor string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["TextColor"] = true
	h.TextColor = &textColor
}

func (h *Header) SetTextColorNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["TextColor"] = true
	h.TextColor = nil
}
func (h Header) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if h.touched["BackgroundColor"] && h.BackgroundColor == nil {
		data["backgroundColor"] = nil
	} else if h.BackgroundColor != nil {
		data["backgroundColor"] = h.BackgroundColor
	}

	if h.touched["BorderRadius"] && h.BorderRadius == nil {
		data["borderRadius"] = nil
	} else if h.BorderRadius != nil {
		data["borderRadius"] = h.BorderRadius
	}

	if h.touched["FontFamily"] && h.FontFamily == nil {
		data["fontFamily"] = nil
	} else if h.FontFamily != nil {
		data["fontFamily"] = h.FontFamily
	}

	if h.touched["FontSize"] && h.FontSize == nil {
		data["fontSize"] = nil
	} else if h.FontSize != nil {
		data["fontSize"] = h.FontSize
	}

	if h.touched["TextColor"] && h.TextColor == nil {
		data["textColor"] = nil
	} else if h.TextColor != nil {
		data["textColor"] = h.TextColor
	}

	return json.Marshal(data)
}

type Icon struct {
	BorderColor *string `json:"borderColor,omitempty" required:"true"`
	Width       *string `json:"width,omitempty" required:"true"`
	touched     map[string]bool
}

func (i *Icon) GetBorderColor() *string {
	if i == nil {
		return nil
	}
	return i.BorderColor
}

func (i *Icon) SetBorderColor(borderColor string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["BorderColor"] = true
	i.BorderColor = &borderColor
}

func (i *Icon) SetBorderColorNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["BorderColor"] = true
	i.BorderColor = nil
}

func (i *Icon) GetWidth() *string {
	if i == nil {
		return nil
	}
	return i.Width
}

func (i *Icon) SetWidth(width string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Width"] = true
	i.Width = &width
}

func (i *Icon) SetWidthNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Width"] = true
	i.Width = nil
}
func (i Icon) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if i.touched["BorderColor"] && i.BorderColor == nil {
		data["borderColor"] = nil
	} else if i.BorderColor != nil {
		data["borderColor"] = i.BorderColor
	}

	if i.touched["Width"] && i.Width == nil {
		data["width"] = nil
	} else if i.Width != nil {
		data["width"] = i.Width
	}

	return json.Marshal(data)
}

type Notification struct {
	Default_ *Default_ `json:"default,omitempty" required:"true"`
	Unread   *Unread   `json:"unread,omitempty" required:"true"`
	Unseen   *Unseen   `json:"unseen,omitempty" required:"true"`
	touched  map[string]bool
}

func (n *Notification) GetDefault_() *Default_ {
	if n == nil {
		return nil
	}
	return n.Default_
}

func (n *Notification) SetDefault_(default_ Default_) {
	if n.touched == nil {
		n.touched = map[string]bool{}
	}
	n.touched["Default_"] = true
	n.Default_ = &default_
}

func (n *Notification) SetDefault_Nil() {
	if n.touched == nil {
		n.touched = map[string]bool{}
	}
	n.touched["Default_"] = true
	n.Default_ = nil
}

func (n *Notification) GetUnread() *Unread {
	if n == nil {
		return nil
	}
	return n.Unread
}

func (n *Notification) SetUnread(unread Unread) {
	if n.touched == nil {
		n.touched = map[string]bool{}
	}
	n.touched["Unread"] = true
	n.Unread = &unread
}

func (n *Notification) SetUnreadNil() {
	if n.touched == nil {
		n.touched = map[string]bool{}
	}
	n.touched["Unread"] = true
	n.Unread = nil
}

func (n *Notification) GetUnseen() *Unseen {
	if n == nil {
		return nil
	}
	return n.Unseen
}

func (n *Notification) SetUnseen(unseen Unseen) {
	if n.touched == nil {
		n.touched = map[string]bool{}
	}
	n.touched["Unseen"] = true
	n.Unseen = &unseen
}

func (n *Notification) SetUnseenNil() {
	if n.touched == nil {
		n.touched = map[string]bool{}
	}
	n.touched["Unseen"] = true
	n.Unseen = nil
}
func (n Notification) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if n.touched["Default_"] && n.Default_ == nil {
		data["default"] = nil
	} else if n.Default_ != nil {
		data["default"] = n.Default_
	}

	if n.touched["Unread"] && n.Unread == nil {
		data["unread"] = nil
	} else if n.Unread != nil {
		data["unread"] = n.Unread
	}

	if n.touched["Unseen"] && n.Unseen == nil {
		data["unseen"] = nil
	} else if n.Unseen != nil {
		data["unseen"] = n.Unseen
	}

	return json.Marshal(data)
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
	touched         map[string]bool
}

func (d *Default_) GetBackgroundColor() *string {
	if d == nil {
		return nil
	}
	return d.BackgroundColor
}

func (d *Default_) SetBackgroundColor(backgroundColor string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["BackgroundColor"] = true
	d.BackgroundColor = &backgroundColor
}

func (d *Default_) SetBackgroundColorNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["BackgroundColor"] = true
	d.BackgroundColor = nil
}

func (d *Default_) GetBorderRadius() *string {
	if d == nil {
		return nil
	}
	return d.BorderRadius
}

func (d *Default_) SetBorderRadius(borderRadius string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["BorderRadius"] = true
	d.BorderRadius = &borderRadius
}

func (d *Default_) SetBorderRadiusNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["BorderRadius"] = true
	d.BorderRadius = nil
}

func (d *Default_) GetFontFamily() *string {
	if d == nil {
		return nil
	}
	return d.FontFamily
}

func (d *Default_) SetFontFamily(fontFamily string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["FontFamily"] = true
	d.FontFamily = &fontFamily
}

func (d *Default_) SetFontFamilyNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["FontFamily"] = true
	d.FontFamily = nil
}

func (d *Default_) GetFontSize() *string {
	if d == nil {
		return nil
	}
	return d.FontSize
}

func (d *Default_) SetFontSize(fontSize string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["FontSize"] = true
	d.FontSize = &fontSize
}

func (d *Default_) SetFontSizeNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["FontSize"] = true
	d.FontSize = nil
}

func (d *Default_) GetHover() *DefaultHover {
	if d == nil {
		return nil
	}
	return d.Hover
}

func (d *Default_) SetHover(hover DefaultHover) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Hover"] = true
	d.Hover = &hover
}

func (d *Default_) SetHoverNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Hover"] = true
	d.Hover = nil
}

func (d *Default_) GetMargin() *string {
	if d == nil {
		return nil
	}
	return d.Margin
}

func (d *Default_) SetMargin(margin string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Margin"] = true
	d.Margin = &margin
}

func (d *Default_) SetMarginNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Margin"] = true
	d.Margin = nil
}

func (d *Default_) GetState() *DefaultState {
	if d == nil {
		return nil
	}
	return d.State
}

func (d *Default_) SetState(state DefaultState) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["State"] = true
	d.State = &state
}

func (d *Default_) SetStateNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["State"] = true
	d.State = nil
}

func (d *Default_) GetTextColor() *string {
	if d == nil {
		return nil
	}
	return d.TextColor
}

func (d *Default_) SetTextColor(textColor string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["TextColor"] = true
	d.TextColor = &textColor
}

func (d *Default_) SetTextColorNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["TextColor"] = true
	d.TextColor = nil
}
func (d Default_) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["BackgroundColor"] && d.BackgroundColor == nil {
		data["backgroundColor"] = nil
	} else if d.BackgroundColor != nil {
		data["backgroundColor"] = d.BackgroundColor
	}

	if d.touched["BorderRadius"] && d.BorderRadius == nil {
		data["borderRadius"] = nil
	} else if d.BorderRadius != nil {
		data["borderRadius"] = d.BorderRadius
	}

	if d.touched["FontFamily"] && d.FontFamily == nil {
		data["fontFamily"] = nil
	} else if d.FontFamily != nil {
		data["fontFamily"] = d.FontFamily
	}

	if d.touched["FontSize"] && d.FontSize == nil {
		data["fontSize"] = nil
	} else if d.FontSize != nil {
		data["fontSize"] = d.FontSize
	}

	if d.touched["Hover"] && d.Hover == nil {
		data["hover"] = nil
	} else if d.Hover != nil {
		data["hover"] = d.Hover
	}

	if d.touched["Margin"] && d.Margin == nil {
		data["margin"] = nil
	} else if d.Margin != nil {
		data["margin"] = d.Margin
	}

	if d.touched["State"] && d.State == nil {
		data["state"] = nil
	} else if d.State != nil {
		data["state"] = d.State
	}

	if d.touched["TextColor"] && d.TextColor == nil {
		data["textColor"] = nil
	} else if d.TextColor != nil {
		data["textColor"] = d.TextColor
	}

	return json.Marshal(data)
}

type DefaultHover struct {
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	touched         map[string]bool
}

func (d *DefaultHover) GetBackgroundColor() *string {
	if d == nil {
		return nil
	}
	return d.BackgroundColor
}

func (d *DefaultHover) SetBackgroundColor(backgroundColor string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["BackgroundColor"] = true
	d.BackgroundColor = &backgroundColor
}

func (d *DefaultHover) SetBackgroundColorNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["BackgroundColor"] = true
	d.BackgroundColor = nil
}
func (d DefaultHover) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["BackgroundColor"] && d.BackgroundColor == nil {
		data["backgroundColor"] = nil
	} else if d.BackgroundColor != nil {
		data["backgroundColor"] = d.BackgroundColor
	}

	return json.Marshal(data)
}

type DefaultState struct {
	Color   *string `json:"color,omitempty" required:"true"`
	touched map[string]bool
}

func (d *DefaultState) GetColor() *string {
	if d == nil {
		return nil
	}
	return d.Color
}

func (d *DefaultState) SetColor(color string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Color"] = true
	d.Color = &color
}

func (d *DefaultState) SetColorNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Color"] = true
	d.Color = nil
}
func (d DefaultState) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["Color"] && d.Color == nil {
		data["color"] = nil
	} else if d.Color != nil {
		data["color"] = d.Color
	}

	return json.Marshal(data)
}

type Unread struct {
	BackgroundColor *string      `json:"backgroundColor,omitempty" required:"true"`
	Hover           *UnreadHover `json:"hover,omitempty"`
	State           *UnreadState `json:"state,omitempty"`
	TextColor       *string      `json:"textColor,omitempty" required:"true"`
	touched         map[string]bool
}

func (u *Unread) GetBackgroundColor() *string {
	if u == nil {
		return nil
	}
	return u.BackgroundColor
}

func (u *Unread) SetBackgroundColor(backgroundColor string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["BackgroundColor"] = true
	u.BackgroundColor = &backgroundColor
}

func (u *Unread) SetBackgroundColorNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["BackgroundColor"] = true
	u.BackgroundColor = nil
}

func (u *Unread) GetHover() *UnreadHover {
	if u == nil {
		return nil
	}
	return u.Hover
}

func (u *Unread) SetHover(hover UnreadHover) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Hover"] = true
	u.Hover = &hover
}

func (u *Unread) SetHoverNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Hover"] = true
	u.Hover = nil
}

func (u *Unread) GetState() *UnreadState {
	if u == nil {
		return nil
	}
	return u.State
}

func (u *Unread) SetState(state UnreadState) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["State"] = true
	u.State = &state
}

func (u *Unread) SetStateNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["State"] = true
	u.State = nil
}

func (u *Unread) GetTextColor() *string {
	if u == nil {
		return nil
	}
	return u.TextColor
}

func (u *Unread) SetTextColor(textColor string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["TextColor"] = true
	u.TextColor = &textColor
}

func (u *Unread) SetTextColorNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["TextColor"] = true
	u.TextColor = nil
}
func (u Unread) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["BackgroundColor"] && u.BackgroundColor == nil {
		data["backgroundColor"] = nil
	} else if u.BackgroundColor != nil {
		data["backgroundColor"] = u.BackgroundColor
	}

	if u.touched["Hover"] && u.Hover == nil {
		data["hover"] = nil
	} else if u.Hover != nil {
		data["hover"] = u.Hover
	}

	if u.touched["State"] && u.State == nil {
		data["state"] = nil
	} else if u.State != nil {
		data["state"] = u.State
	}

	if u.touched["TextColor"] && u.TextColor == nil {
		data["textColor"] = nil
	} else if u.TextColor != nil {
		data["textColor"] = u.TextColor
	}

	return json.Marshal(data)
}

type UnreadHover struct {
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	touched         map[string]bool
}

func (u *UnreadHover) GetBackgroundColor() *string {
	if u == nil {
		return nil
	}
	return u.BackgroundColor
}

func (u *UnreadHover) SetBackgroundColor(backgroundColor string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["BackgroundColor"] = true
	u.BackgroundColor = &backgroundColor
}

func (u *UnreadHover) SetBackgroundColorNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["BackgroundColor"] = true
	u.BackgroundColor = nil
}
func (u UnreadHover) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["BackgroundColor"] && u.BackgroundColor == nil {
		data["backgroundColor"] = nil
	} else if u.BackgroundColor != nil {
		data["backgroundColor"] = u.BackgroundColor
	}

	return json.Marshal(data)
}

type UnreadState struct {
	Color   *string `json:"color,omitempty" required:"true"`
	touched map[string]bool
}

func (u *UnreadState) GetColor() *string {
	if u == nil {
		return nil
	}
	return u.Color
}

func (u *UnreadState) SetColor(color string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Color"] = true
	u.Color = &color
}

func (u *UnreadState) SetColorNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Color"] = true
	u.Color = nil
}
func (u UnreadState) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["Color"] && u.Color == nil {
		data["color"] = nil
	} else if u.Color != nil {
		data["color"] = u.Color
	}

	return json.Marshal(data)
}

type Unseen struct {
	BackgroundColor *string      `json:"backgroundColor,omitempty" required:"true"`
	Hover           *UnseenHover `json:"hover,omitempty"`
	State           *UnseenState `json:"state,omitempty"`
	TextColor       *string      `json:"textColor,omitempty" required:"true"`
	touched         map[string]bool
}

func (u *Unseen) GetBackgroundColor() *string {
	if u == nil {
		return nil
	}
	return u.BackgroundColor
}

func (u *Unseen) SetBackgroundColor(backgroundColor string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["BackgroundColor"] = true
	u.BackgroundColor = &backgroundColor
}

func (u *Unseen) SetBackgroundColorNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["BackgroundColor"] = true
	u.BackgroundColor = nil
}

func (u *Unseen) GetHover() *UnseenHover {
	if u == nil {
		return nil
	}
	return u.Hover
}

func (u *Unseen) SetHover(hover UnseenHover) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Hover"] = true
	u.Hover = &hover
}

func (u *Unseen) SetHoverNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Hover"] = true
	u.Hover = nil
}

func (u *Unseen) GetState() *UnseenState {
	if u == nil {
		return nil
	}
	return u.State
}

func (u *Unseen) SetState(state UnseenState) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["State"] = true
	u.State = &state
}

func (u *Unseen) SetStateNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["State"] = true
	u.State = nil
}

func (u *Unseen) GetTextColor() *string {
	if u == nil {
		return nil
	}
	return u.TextColor
}

func (u *Unseen) SetTextColor(textColor string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["TextColor"] = true
	u.TextColor = &textColor
}

func (u *Unseen) SetTextColorNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["TextColor"] = true
	u.TextColor = nil
}
func (u Unseen) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["BackgroundColor"] && u.BackgroundColor == nil {
		data["backgroundColor"] = nil
	} else if u.BackgroundColor != nil {
		data["backgroundColor"] = u.BackgroundColor
	}

	if u.touched["Hover"] && u.Hover == nil {
		data["hover"] = nil
	} else if u.Hover != nil {
		data["hover"] = u.Hover
	}

	if u.touched["State"] && u.State == nil {
		data["state"] = nil
	} else if u.State != nil {
		data["state"] = u.State
	}

	if u.touched["TextColor"] && u.TextColor == nil {
		data["textColor"] = nil
	} else if u.TextColor != nil {
		data["textColor"] = u.TextColor
	}

	return json.Marshal(data)
}

type UnseenHover struct {
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	touched         map[string]bool
}

func (u *UnseenHover) GetBackgroundColor() *string {
	if u == nil {
		return nil
	}
	return u.BackgroundColor
}

func (u *UnseenHover) SetBackgroundColor(backgroundColor string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["BackgroundColor"] = true
	u.BackgroundColor = &backgroundColor
}

func (u *UnseenHover) SetBackgroundColorNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["BackgroundColor"] = true
	u.BackgroundColor = nil
}
func (u UnseenHover) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["BackgroundColor"] && u.BackgroundColor == nil {
		data["backgroundColor"] = nil
	} else if u.BackgroundColor != nil {
		data["backgroundColor"] = u.BackgroundColor
	}

	return json.Marshal(data)
}

type UnseenState struct {
	Color   *string `json:"color,omitempty" required:"true"`
	touched map[string]bool
}

func (u *UnseenState) GetColor() *string {
	if u == nil {
		return nil
	}
	return u.Color
}

func (u *UnseenState) SetColor(color string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Color"] = true
	u.Color = &color
}

func (u *UnseenState) SetColorNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Color"] = true
	u.Color = nil
}
func (u UnseenState) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["Color"] && u.Color == nil {
		data["color"] = nil
	} else if u.Color != nil {
		data["color"] = u.Color
	}

	return json.Marshal(data)
}

type UnseenBadge struct {
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	touched         map[string]bool
}

func (u *UnseenBadge) GetBackgroundColor() *string {
	if u == nil {
		return nil
	}
	return u.BackgroundColor
}

func (u *UnseenBadge) SetBackgroundColor(backgroundColor string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["BackgroundColor"] = true
	u.BackgroundColor = &backgroundColor
}

func (u *UnseenBadge) SetBackgroundColorNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["BackgroundColor"] = true
	u.BackgroundColor = nil
}
func (u UnseenBadge) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["BackgroundColor"] && u.BackgroundColor == nil {
		data["backgroundColor"] = nil
	} else if u.BackgroundColor != nil {
		data["backgroundColor"] = u.BackgroundColor
	}

	return json.Marshal(data)
}
