# ttsmaker
ttsmaker is a Text-to-Speech library implemented using the TTSMaker API.  

API documents is [here](https://ttsmaker.com/developer-api-docs).  

## Example  

New client with you token:  

```go
client := ttsmaker.NewClient("ttsmaker_demo_token")
```  

### Get All Voice List  

```go
listResp, err := client.GetVoiceList(context.Background())
```  

Or, get voice list with languages:  

```go
listResp, err := client.GetVoiceListWithLang(context.Background(), "en")  
```  

> support language: ["en","zh","es","ja","ko","de","fr","it","ru","pt","tr","ms","th","vi","id","he"]  

### Get Yout Token Status  

```go
statusResp, err := client.GetTokenStatus(context.Background())
```  

### Create a TTS Order  

```go
orderRequest := &ttsmaker.OrderRequest{
	Token:                  "ttsmaker_demo_token",
	Text:                   "hello, world!",
	VoiceID:                "778",
	AudioFormat:            "mp3",
	AudioSpeed:             1.0,
	AudioVolume:            0,
	TextParagraphPauseTime: 0,
}

orderResp, err := client.CreateOrder(context.Background(), orderRequest)
```  

