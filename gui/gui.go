package gui

import (
	"encoding/json"
	"fmt"
	"helperGPT/api"
	"helperGPT/gpt"
	"helperGPT/gpt/scenario"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var tabs *container.AppTabs
var tabMap map[string]*container.TabItem
var usedModel string

func LoadGui() {
	a := app.New()
	a.SetIcon(theme.QuestionIcon())
	w := a.NewWindow("helperGPT")

	tabs = &container.AppTabs{}
	tabMap = make(map[string]*container.TabItem)

	btnGrid := container.New(
		layout.NewGridWrapLayout(
			fyne.NewSize(200, 40)),
		newScenarioButton(scenario.PersonalAssistant),
		newScenarioButton(scenario.Chatbot),
		newScenarioButton(scenario.Programmer),
		newScenarioButton(scenario.PolishEnglishTranslator),
		newScenarioButton(scenario.Teacher),
	)

	appGrid := container.New(layout.NewVBoxLayout(),
		newGptSelectList(),
		// NewSlider())
		btnGrid,
	)

	w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		if event.Name == fyne.KeyReturn {
			appGrid.Objects[3].(*widget.Form).OnSubmit()
		}
	})

	homeContainer := container.NewTabItemWithIcon("Home", theme.HomeIcon(), appGrid)

	tabs.Append(homeContainer)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(748, 658))
	w.ShowAndRun()
}

func newGptSelectList() *widget.Select {
	gptSelectList := widget.NewSelect(
		[]string{
			gpt.GPT3_5Turbo,
			gpt.GPT4,
		},
		func(model string) {
			usedModel = model
		},
	)
	gptSelectList.SetSelectedIndex(0)
	gptSelectList.Resize(fyne.NewSize(120, 40))

	return gptSelectList
}

func newScenarioButton(scenarioName string) *widget.Button {

	button := widget.NewButton(scenarioName, func() {
		if tabMap[scenarioIdGenerator(scenarioName)] == nil {
			gptScenario := *gpt.NewGpt(usedModel, scenarioName)
			inputForm := newConversationForm(&gptScenario)

			tabMap[scenarioIdGenerator(scenarioName)] = (container.NewTabItem(scenarioIdGenerator(scenarioName), inputForm))
			tabs.Append(tabMap[scenarioIdGenerator(scenarioName)])

			inputForm.OnCancel = func() {
				tabs.Remove(tabMap[scenarioIdGenerator(gptScenario.Conversation.Scenario, gptScenario.Conversation.Model)])
				delete(tabMap, scenarioIdGenerator(gptScenario.Conversation.Scenario, gptScenario.Conversation.Model))
			}
		}
		tabs.Select(tabMap[scenarioIdGenerator(scenarioName)])
	})
	return button
}

func newConversationForm(gptScenario *gpt.GptScenario) *widget.Form {

	input := widget.NewMultiLineEntry()
	input.SetMinRowsVisible(5)

	response := widget.NewMultiLineEntry()
	response.SetMinRowsVisible(20)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Response", Widget: response},
			{Text: "Input", Widget: input}},
		SubmitText: "Send",
		CancelText: "Close",
		OnSubmit: func() {

			log.Println("Input:", input.Text)

			gptScenario.Conversation.AddMessage(gpt.NewMessage(gpt.User, input.Text))

			input.SetText("")
			response.SetText(gptScenario.Conversation.PrintConversation())

			resp, err := api.GetResponse(*gptScenario.Conversation)
			if err != nil {
				log.Println(err)
			}

			var result interface{}
			err = json.Unmarshal([]byte(resp), &result)
			if err != nil {
				log.Println(err)
			}

			choices := result.(map[string]interface{})["choices"].([]interface{})[0].(map[string]interface{})["message"]

			out, err := json.Marshal(choices)
			if err != nil {
				log.Println(err)
			}

			var responseMessage gpt.Message
			json.Unmarshal(out, &responseMessage)

			gptScenario.Conversation.AddMessage(responseMessage)

			log.Println("Response:", responseMessage.Content)

			response.SetText(gptScenario.Conversation.PrintConversation())
		},
	}

	return form
}

func scenarioIdGenerator(scenarioName string, modelName ...string) string {
	if len(modelName) > 0 {
		usedModel = modelName[0]
	}

	return fmt.Sprintf("%s (%s)", scenarioName, usedModel)
}

// func NewSliderWrapper() *fyne.Container {
// 	slider := NewSlider()
// 	sliderWrapper := container.New(layout.NewHBoxLayout(),
// 		widget.NewLabel("Temperature"),
// 		slider,
// 		// widget.NewLabel(value),
// 	)
// 	return sliderWrapper
// }

// func NewSlider() *widget.Slider {
// 	slider := widget.NewSlider(0.0, 1.0)
// 	slider.Step = 0.1
// 	slider.OnChanged = func(value float64) {
// 		gpt.SetTemperature(float32(value))
// 	}

// 	return slider
// }
