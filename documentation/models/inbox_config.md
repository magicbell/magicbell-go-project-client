# InboxConfig

**Properties**

| Name   | Type                | Required | Description |
| :----- | :------------------ | :------- | :---------- |
| Images | integrations.Images | ✅       |             |
| Locale | string              | ✅       |             |
| Theme  | integrations.Theme  | ✅       |             |

# Images

**Properties**

| Name          | Type   | Required | Description |
| :------------ | :----- | :------- | :---------- |
| EmptyInboxUrl | string | ✅       |             |

# Theme

**Properties**

| Name         | Type                      | Required | Description |
| :----------- | :------------------------ | :------- | :---------- |
| Banner       | integrations.Banner       | ❌       |             |
| Dialog       | integrations.Dialog       | ❌       |             |
| Footer       | integrations.Footer       | ❌       |             |
| Header       | integrations.Header       | ❌       |             |
| Icon         | integrations.Icon         | ❌       |             |
| Notification | integrations.Notification | ❌       |             |
| UnseenBadge  | integrations.UnseenBadge  | ❌       |             |

# Banner

**Properties**

| Name              | Type    | Required | Description |
| :---------------- | :------ | :------- | :---------- |
| BackgroundColor   | string  | ✅       |             |
| FontSize          | string  | ✅       |             |
| TextColor         | string  | ✅       |             |
| BackgroundOpacity | float64 | ❌       |             |

# Dialog

**Properties**

| Name            | Type   | Required | Description |
| :-------------- | :----- | :------- | :---------- |
| AccentColor     | string | ✅       |             |
| BackgroundColor | string | ✅       |             |
| TextColor       | string | ✅       |             |

# Footer

**Properties**

| Name            | Type   | Required | Description |
| :-------------- | :----- | :------- | :---------- |
| BackgroundColor | string | ✅       |             |
| BorderRadius    | string | ✅       |             |
| FontSize        | string | ✅       |             |
| TextColor       | string | ✅       |             |

# Header

**Properties**

| Name            | Type   | Required | Description |
| :-------------- | :----- | :------- | :---------- |
| BackgroundColor | string | ✅       |             |
| BorderRadius    | string | ✅       |             |
| FontFamily      | string | ✅       |             |
| FontSize        | string | ✅       |             |
| TextColor       | string | ✅       |             |

# Icon

**Properties**

| Name        | Type   | Required | Description |
| :---------- | :----- | :------- | :---------- |
| BorderColor | string | ✅       |             |
| Width       | string | ✅       |             |

# Notification

**Properties**

| Name      | Type                   | Required | Description |
| :-------- | :--------------------- | :------- | :---------- |
| Default\_ | integrations.Default\_ | ✅       |             |
| Unread    | integrations.Unread    | ✅       |             |
| Unseen    | integrations.Unseen    | ✅       |             |

# Default\_

**Properties**

| Name            | Type                      | Required | Description |
| :-------------- | :------------------------ | :------- | :---------- |
| BackgroundColor | string                    | ✅       |             |
| BorderRadius    | string                    | ✅       |             |
| FontFamily      | string                    | ✅       |             |
| FontSize        | string                    | ✅       |             |
| Margin          | string                    | ✅       |             |
| TextColor       | string                    | ✅       |             |
| Hover           | integrations.DefaultHover | ❌       |             |
| State           | integrations.DefaultState | ❌       |             |

# DefaultHover

**Properties**

| Name            | Type   | Required | Description |
| :-------------- | :----- | :------- | :---------- |
| BackgroundColor | string | ✅       |             |

# DefaultState

**Properties**

| Name  | Type   | Required | Description |
| :---- | :----- | :------- | :---------- |
| Color | string | ✅       |             |

# Unread

**Properties**

| Name            | Type                     | Required | Description |
| :-------------- | :----------------------- | :------- | :---------- |
| BackgroundColor | string                   | ✅       |             |
| TextColor       | string                   | ✅       |             |
| Hover           | integrations.UnreadHover | ❌       |             |
| State           | integrations.UnreadState | ❌       |             |

# UnreadHover

**Properties**

| Name            | Type   | Required | Description |
| :-------------- | :----- | :------- | :---------- |
| BackgroundColor | string | ✅       |             |

# UnreadState

**Properties**

| Name  | Type   | Required | Description |
| :---- | :----- | :------- | :---------- |
| Color | string | ✅       |             |

# Unseen

**Properties**

| Name            | Type                     | Required | Description |
| :-------------- | :----------------------- | :------- | :---------- |
| BackgroundColor | string                   | ✅       |             |
| TextColor       | string                   | ✅       |             |
| Hover           | integrations.UnseenHover | ❌       |             |
| State           | integrations.UnseenState | ❌       |             |

# UnseenHover

**Properties**

| Name            | Type   | Required | Description |
| :-------------- | :----- | :------- | :---------- |
| BackgroundColor | string | ✅       |             |

# UnseenState

**Properties**

| Name  | Type   | Required | Description |
| :---- | :----- | :------- | :---------- |
| Color | string | ✅       |             |

# UnseenBadge

**Properties**

| Name            | Type   | Required | Description |
| :-------------- | :----- | :------- | :---------- |
| BackgroundColor | string | ✅       |             |
