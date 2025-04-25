package ttsmaker

// ListResponse represents the response from the voice list API
type ListResponse struct {
	ErrorCode           int                  `json:"error_code"`
	ErrorSummary        string               `json:"error_summary"`
	Msg                 string               `json:"msg"`
	UserEmail           string               `json:"user_email"`
	Language            string               `json:"language"`
	SupportLanguageList []string             `json:"support_language_list"`
	VoicesIDList        []int                `json:"voices_id_list"`
	VoicesCount         int                  `json:"voices_count"`
	VoicesDetailedList  []VoicesDetailedList `json:"voices_detailed_list"`
}

// VoicesDetailedList represents the detailed information of a voice
type VoicesDetailedList struct {
	ID                    int      `json:"id"`
	Name                  string   `json:"name"`
	Language              string   `json:"language"`
	AudioSampleFileURL    string   `json:"audio_sample_file_url"`
	TextCharactersLimit   int      `json:"text_characters_limit"`
	IsSupportEmotion      bool     `json:"is_support_emotion,omitempty"`
	SupportEmotionKeyList []string `json:"support_emotion_key_list,omitempty"`
}

// OrderRequest represents the request parameters for creating a TTS order
type OrderRequest struct {
	APIKey                 string  `json:"api_key"`
	Text                   string  `json:"text"`
	VoiceID                int     `json:"voice_id"`
	AudioFormat            string  `json:"audio_format"`
	AudioSpeed             float64 `json:"audio_speed,omitempty"`
	AudioVolume            int     `json:"audio_volume,omitempty"`
	AudioPitch             float64 `json:"audio_pitch,omitempty"`
	AudioHighQuality       int     `json:"audio_high_quality,omitempty"`
	TextParagraphPauseTime int     `json:"text_paragraph_pause_time,omitempty"`
	EmotionStyleKey        string  `json:"emotion_style_key,omitempty"`
	EmotionIntensity       float64 `json:"emotion_intensity,omitempty"`
}

// OrderResponse represents the response from the TTS API
type OrderResponse struct {
	ErrorCode                    int           `json:"error_code"`
	ErrorSummary                 string        `json:"error_summary"`
	Msg                          string        `json:"msg"`
	UserEmail                    string        `json:"user_email"`
	TTSTaskCharactersCount       int           `json:"tts_task_characters_count"`
	AudioDownloadURL             string        `json:"audio_download_url"`
	AudioDownloadBackupURL       string        `json:"audio_download_backup_url"`
	AudioFileFormat              string        `json:"audio_file_format"`
	CurrentTimestamp             int64         `json:"current_timestamp"`
	AudioFileExpirationTimestamp int64         `json:"audio_file_expiration_timestamp"`
	IsDemoKey                    bool          `json:"is_demo_key"`
	AccountStatus                AccountStatus `json:"account_status"`
}

// AccountStatus represents the user's account status and quota information
type AccountStatus struct {
	QuotaCharacters                int    `json:"quota_characters"`
	CharactersUsed                 int    `json:"characters_used"`
	AvailableQuota                 int    `json:"available_quota"`
	SubscriptionPeriod             string `json:"subscription_period"`
	SubscriptionNextResetTimestamp int64  `json:"subscription_next_reset_timestamp"`
}

// StatusResponse represents the response from the API key/account status check
type StatusResponse struct {
	ErrorCode        int           `json:"error_code"`
	ErrorSummary     string        `json:"error_summary"`
	Msg              string        `json:"msg"`
	UserEmail        string        `json:"user_email"`
	CurrentTimestamp int64         `json:"current_timestamp"`
	IsDemoKey        bool          `json:"is_demo_key"`
	AccountStatus    AccountStatus `json:"account_status"`
}
