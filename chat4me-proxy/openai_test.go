package main

import (
	"testing"
)

const (
	justYouConversation = `You:right
You:As long as we have a little, that's fine. We can refine it later
You:Sure
You:Alright
You:Are you guys in the lab room?
You:the chat4me-util library
You:She's not in today so you don't have to come. I forgot until I already got to class so I'm just working on the util functions
You:Do you still have the original document file that you used to create the pdf?
You:I added the state diagram to the document
You:Did we need to have the updated class diagram part of the submission or is that just for resubmission of the structural model assignment?
You:I don't know how but ArgoUML took a dump when it saved the latest UseCaseUML1.zargo so now it's unopenable :/
You:☝️
You:206
You:Since chat4me.org is the cheapest option for the domain, I'm going with that
You:It generates a jar file that you'll be able to import into the main app via Android Studio
You:I've started work on the app utility stuff that I mentioned, though it doesn't do anything particularly interesting yet
https://github.com/The-Chatastic-4-CSCD-350/chat4me-util
You:I submitted the structural model assignment
You:Updated the use case diagram and added another for the auto reply if driving use case
You:nvm, disregard that, I think it might be easier to just share changes by uploading it directly to Discord, the file size is pretty small
You:I shared the Argo UML file with you guys on Google Drive, here's a link
You:Continuing the conversation, to help with the coding stuff, I would look into how to encode and decode JSON, and how to send HTTPS POST requests with form data and custom HTTP headers in Java`

	twoPersonConversation = `Friend:Hello!
You:Hi, how's it going?
Friend:fine, and you?
You:I'm doing alright.
Friend:How is the project going?
You:It's going alright, I just posted changes. Are you still working on your part?
Friend:Yes. What else do we need?`

	threePersonConversation = `Alice:Hello everyone!
Bob:Hello, how is everyone?
You:Hello everyone. I'm pretty good today.
Alice:Good, how about you?
Bob:I'm doing pretty alright.
You:How is the project going?
Bob:Unfortunately, we're running a bit behind schedule.`
)

func TestCompletion(t *testing.T) {
	initConfig()
	initOpenAI()
	resp, err := doCompletion(justYouConversation)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("Generated message for 'single person' (just You) conversation: %q", resp.Choices[0].Text)
	resp, err = doCompletion(twoPersonConversation)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("Generated message for two person conversation: %q", resp.Choices[0].Text)
	resp, err = doCompletion(threePersonConversation)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("Generated message for three person conversation: %q", resp.Choices[0].Text)
}
