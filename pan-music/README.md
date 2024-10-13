# Pan-Music

表设计:

设计一个音乐网站的后台数据库时，通常需要考虑以下几个主要表及其字段：

## 1. **用户表 (Users)**

- **user_id**: 主键，唯一标识用户
- **username**: 用户名
- **email**: 电子邮件
- **password_hash**: 密码哈希
- **created_at**: 创建时间
- **updated_at**: 更新时间

### 2. **音乐表 (Songs)**

- **song_id**: 主键，唯一标识歌曲
- **title**: 歌曲标题
- **artist_id**: 外键，指向艺术家表
- **album_id**: 外键，指向专辑表
- **genre**: 音乐类型
- **duration**: 时长（以秒为单位）
- **release_date**: 发行日期
- **lyrics**: 歌词文本
- **created_at**: 创建时间
- **updated_at**: 更新时间

### 3. **艺术家表 (Artists)**

- **artist_id**: 主键，唯一标识艺术家
- **name**: 艺术家名称
- **bio**: 简介
- **created_at**: 创建时间
- **updated_at**: 更新时间

### 4. **专辑表 (Albums)**

- **album_id**: 主键，唯一标识专辑
- **title**: 专辑标题
- **artist_id**: 外键，指向艺术家表
- **release_date**: 发行日期
- **cover_image**: 封面图片URL
- **created_at**: 创建时间
- **updated_at**: 更新时间

### 5. **播放列表表 (Playlists)**

- **playlist_id**: 主键，唯一标识播放列表
- **user_id**: 外键，指向用户表
- **title**: 播放列表标题
- **description**: 描述
- **created_at**: 创建时间
- **updated_at**: 更新时间

### 6. **播放列表歌曲表 (Playlist_Songs)**

- **playlist_id**: 外键，指向播放列表表
- **song_id**: 外键，指向音乐表
- **added_at**: 添加时间

### 7. **评论表 (Comments)**

- **comment_id**: 主键，唯一标识评论
- **user_id**: 外键，指向用户表
- **song_id**: 外键，指向音乐表
- **content**: 评论内容
- **created_at**: 创建时间

### 8. **喜欢表 (Likes)**

- **like_id**: 主键，唯一标识喜欢记录
- **user_id**: 外键，指向用户表
- **song_id**: 外键，指向音乐表
- **created_at**: 创建时间

### 9. **用户活动表 (User_Actions)**

- **action_id**: 主键，唯一标识活动记录
- **user_id**: 外键，指向用户表
- **action_type**: 活动类型（例如，播放、喜欢、评论等）
- **song_id**: 外键，指向音乐表（如果适用）
- **created_at**: 创建时间

这种设计可以为网站的核心功能提供支持，并能方便地进行扩展。你有特定的需求或功能想要添加吗？
针对这些表，可以设计以下接口：

## 1. **用户接口**

- **注册用户**: POST `/api/users/register`
- **用户登录**: POST `/api/users/login`
- **获取用户信息**: GET `/api/users/:user_id`
- **更新用户信息**: PUT `/api/users/:user_id`
- **删除用户**: DELETE `/api/users/:user_id`

### 2. **音乐接口**

- **上传歌曲**: POST `/api/songs`
- **获取歌曲列表**: GET `/api/songs`
- **获取单曲信息**: GET `/api/songs/:song_id`
- **更新歌曲信息**: PUT `/api/songs/:song_id`
- **删除歌曲**: DELETE `/api/songs/:song_id`

### 3. **艺术家接口**

- **添加艺术家**: POST `/api/artists`
- **获取艺术家列表**: GET `/api/artists`
- **获取单个艺术家信息**: GET `/api/artists/:artist_id`
- **更新艺术家信息**: PUT `/api/artists/:artist_id`
- **删除艺术家**: DELETE `/api/artists/:artist_id`

### 4. **专辑接口**

- **添加专辑**: POST `/api/albums`
- **获取专辑列表**: GET `/api/albums`
- **获取单个专辑信息**: GET `/api/albums/:album_id`
- **更新专辑信息**: PUT `/api/albums/:album_id`
- **删除专辑**: DELETE `/api/albums/:album_id`

### 5. **播放列表接口**

- **创建播放列表**: POST `/api/playlists`
- **获取用户播放列表**: GET `/api/users/:user_id/playlists`
- **获取播放列表内容**: GET `/api/playlists/:playlist_id`
- **添加歌曲到播放列表**: POST `/api/playlists/:playlist_id/songs`
- **删除播放列表**: DELETE `/api/playlists/:playlist_id`

### 6. **评论接口**

- **添加评论**: POST `/api/comments`
- **获取歌曲评论**: GET `/api/songs/:song_id/comments`
- **删除评论**: DELETE `/api/comments/:comment_id`

### 7. **喜欢接口**

- **喜欢歌曲**: POST `/api/songs/:song_id/like`
- **取消喜欢**: DELETE `/api/songs/:song_id/like`
- **获取用户喜欢的歌曲**: GET `/api/users/:user_id/likes`

### 8. **用户活动接口**

- **记录用户活动**: POST `/api/user-actions`
- **获取用户活动记录**: GET `/api/users/:user_id/actions`

为了增加社区功能，可以考虑以下表及字段：

### 1. **社区表 (Communities)**

- **community_id**: 主键，唯一标识社区
- **name**: 社区名称
- **description**: 描述
- **created_at**: 创建时间
- **updated_at**: 更新时间

### 2. **话题表 (Topics)**

- **topic_id**: 主键，唯一标识话题
- **community_id**: 外键，指向社区表
- **title**: 话题标题
- **content**: 话题内容
- **user_id**: 外键，指向用户表
- **created_at**: 创建时间
- **updated_at**: 更新时间

### 3. **评论表 (Topic_Comments)**

- **comment_id**: 主键，唯一标识评论
- **topic_id**: 外键，指向话题表
- **user_id**: 外键，指向用户表
- **content**: 评论内容
- **created_at**: 创建时间

### 4. **社区成员表 (Community_Members)**

- **member_id**: 主键，唯一标识成员记录
- **community_id**: 外键，指向社区表
- **user_id**: 外键，指向用户表
- **joined_at**: 加入时间

### 接口设计

#### 1. **社区接口**

- **创建社区**: POST `/api/communities`
- **获取社区列表**: GET `/api/communities`
- **获取单个社区信息**: GET `/api/communities/:community_id`
- **更新社区信息**: PUT `/api/communities/:community_id`
- **删除社区**: DELETE `/api/communities/:community_id`

#### 2. **话题接口**

- **创建话题**: POST `/api/communities/:community_id/topics`
- **获取社区话题列表**: GET `/api/communities/:community_id/topics`
- **获取单个话题信息**: GET `/api/topics/:topic_id`
- **更新话题信息**: PUT `/api/topics/:topic_id`
- **删除话题**: DELETE `/api/topics/:topic_id`

#### 3. **评论接口**

- **添加评论**: POST `/api/topics/:topic_id/comments`
- **获取话题评论**: GET `/api/topics/:topic_id/comments`
- **删除评论**: DELETE `/api/comments/:comment_id`

#### 4. **社区成员接口**

- **加入社区**: POST `/api/communities/:community_id/members`
- **获取社区成员列表**: GET `/api/communities/:community_id/members`
- **退出社区**: DELETE `/api/communities/:community_id/members/:user_id`
