package ttsmaker

// for voice list
type ListResponse struct {
	ErrorCode           string               `json:"error_code"`
	Status              string               `json:"status"`
	ErrorDetails        string               `json:"error_details"`
	Token               string               `json:"token"`
	Language            string               `json:"language"`
	APIMaxQps           string               `json:"api_max_qps"`
	SupportLanguageList []string             `json:"support_language_list"`
	VoicesIDList        []int64              `json:"voices_id_list"`
	VoicesDetailedList  []VoicesDetailedList `json:"voices_detailed_list"`
	VoicesCount         int64                `json:"voices_count"`
}

type VoicesDetailedList struct {
	Name         string `json:"name"`
	Language     string `json:"language"`
	SampleMp3URL string `json:"sample_mp3_url"`
	ID           int64  `json:"id"`
	LimitText    int64  `json:"limit_text"`
}

// for create a ttsmaker order request
type OrderRequest struct {
	Token                  string  `json:"token"`
	Text                   string  `json:"text"`
	VoiceID                string  `json:"voice_id"`
	AudioFormat            string  `json:"audio_format"`
	AudioSpeed             float64 `json:"audio_speed"`
	AudioVolume            int64   `json:"audio_volume"`
	TextParagraphPauseTime int64   `json:"text_paragraph_pause_time"`
}

// order response from ttsmaker
type OrderResponse struct {
	ErrorCode           string      `json:"error_code"`
	Status              string      `json:"status"`
	ErrorDetails        string      `json:"error_details"`
	AudioFileURL        string      `json:"audio_file_url"`
	AudioFileType       string      `json:"audio_file_type"`
	TTSElapsedTime      string      `json:"tts_elapsed_time"`
	TokenStatus         TokenStatus `json:"token_status"`
	UnixTimestamp       int64       `json:"unix_timestamp"`
	AudioFileExpireTime int64       `json:"audio_file_expire_time"`
	TTSOrderCharacters  int64       `json:"tts_order_characters"`
}

// token status
type TokenStatus struct {
	CurrentCycleMaxCharacters       int64   `json:"current_cycle_max_characters"`
	CurrentCycleCharactersUsed      int64   `json:"current_cycle_characters_used"`
	CurrentCycleCharactersAvailable int64   `json:"current_cycle_characters_available"`
	RemainingDaysToResetQuota       float64 `json:"remaining_days_to_reset_quota"`
	HistoryCharactersUsed           int64   `json:"history_characters_used"`
}

// get token status from ttsmaker
type StatusResponse struct {
	ErrorCode   string      `json:"error_code"`
	Status      string      `json:"status"`
	Msg         string      `json:"msg"`
	CurrentTime string      `json:"current_time"`
	Token       string      `json:"token"`
	TokenStatus TokenStatus `json:"token_status"`
}
