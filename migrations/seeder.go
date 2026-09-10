package migrations

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	post2 "lingcard-go/dictionaries/post"
	"lingcard-go/dictionaries/role"
	"lingcard-go/models/language"
	"lingcard-go/models/post"
	"lingcard-go/models/user"
	"lingcard-go/models/voice"
	"lingcard-go/models/word"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type DatabaseSeeder struct {
	db *gorm.DB
}

const batchSize = 1000

func NewSeeder(db *gorm.DB) *DatabaseSeeder {
	return &DatabaseSeeder{db: db}
}

func (s *DatabaseSeeder) RunSeeders() error {
	seeders := []struct {
		name string
		fn   func() error
	}{
		{"LanguageSeeder", s.RunLanguageSeeder},
		{"AvailableLanguageSeed", s.RunAvailableLanguageSeed},
		{"UserSeeder", s.RunUserSeeder},
		{"WordSeeder", s.RunWordSeeder},
		{"PostSeeder", s.RunPostSeeder},
		{"VoteSeeder", s.RunVoteSeeder},
	}

	for _, seeder := range seeders {
		if err := seeder.fn(); err != nil {
			return fmt.Errorf("seeder %s failed: %w", seeder.name, err)
		}
	}

	return nil
}

func (s *DatabaseSeeder) RunLanguageSeeder() error {
	languages := []language.Language{
		{Name: "Русский", Code: "ru", IsActive: true},
		{Name: "Қазақша", Code: "kz", IsActive: true},
		{Name: "English", Code: "en", IsActive: true},
		{Name: "Español", Code: "es", IsActive: true},
		{Name: "Français", Code: "fr", IsActive: true},
		{Name: "Deutsch", Code: "de", IsActive: true},
		{Name: "Italiano", Code: "it", IsActive: false},
		{Name: "Português", Code: "pt", IsActive: true},
		{Name: "Nederlands", Code: "nl", IsActive: false},
		{Name: "Polski", Code: "pl", IsActive: false},
		{Name: "Українська", Code: "ua", IsActive: false},
		{Name: "中文", Code: "cn", IsActive: true},
		{Name: "日本語", Code: "jp", IsActive: true},
		{Name: "한국어", Code: "kr", IsActive: true},
		{Name: "العربية", Code: "sa", IsActive: true},
		{Name: "हिन्दी", Code: "in", IsActive: false},
		{Name: "Türkçe", Code: "tr", IsActive: false},
		{Name: "Tiếng Việt", Code: "vn", IsActive: false},
		{Name: "ไทย", Code: "th", IsActive: false},
		{Name: "Bahasa Indonesia", Code: "id", IsActive: false},
		{Name: "Suomi", Code: "fi", IsActive: false},
		{Name: "Svenska", Code: "se", IsActive: false},
		{Name: "Norsk", Code: "no", IsActive: false},
		{Name: "Dansk", Code: "dk", IsActive: false},
		{Name: "Čeština", Code: "cz", IsActive: false},
		{Name: "Magyar", Code: "hu", IsActive: false},
		{Name: "Română", Code: "ro", IsActive: false},
		{Name: "Slovenčina", Code: "sk", IsActive: false},
		{Name: "Ελληνικά", Code: "gr", IsActive: false},
		{Name: "עברית", Code: "il", IsActive: false},
		{Name: "فارسی", Code: "ir", IsActive: false},
		{Name: "اردو", Code: "pk", IsActive: false},
		{Name: "Монгол", Code: "mn", IsActive: false},
		{Name: "ქართული", Code: "ge", IsActive: false},
		{Name: "Հայերեն", Code: "am", IsActive: false},
		{Name: "Azərbaycan dili", Code: "az", IsActive: false},
		{Name: "Oʻzbekcha", Code: "uz", IsActive: false},
		{Name: "Тоҷикӣ", Code: "tj", IsActive: false},
		{Name: "Кыргызча", Code: "kg", IsActive: false},
		{Name: "Latviešu", Code: "lv", IsActive: false},
		{Name: "Lietuvių", Code: "lt", IsActive: false},
		{Name: "Eesti", Code: "ee", IsActive: false},
		{Name: "Shqip", Code: "al", IsActive: false},
		{Name: "Македонски", Code: "mk", IsActive: false},
		{Name: "Српски", Code: "rs", IsActive: false},
		{Name: "Hrvatski", Code: "hr", IsActive: false},
		{Name: "Bosanski", Code: "ba", IsActive: false},
	}

	return s.db.Create(&languages).Error
}
func (s *DatabaseSeeder) RunAvailableLanguageSeed() error {

	availableCodes := map[string][]string{
		"ru": {"kz", "en", "fr", "cn", "de", "es", "jp", "kr", "pt", "sa"},
		"kz": {"ru", "en", "fr", "cn", "de", "es", "jp", "kr", "pt", "sa"},
		"en": {"ru", "kz", "fr", "cn", "de", "es", "jp", "kr", "pt", "sa"},
		"fr": {"ru", "kz", "en", "cn", "de", "es", "jp", "kr", "pt", "sa"},
		"cn": {"ru", "kz", "en", "fr", "de", "es", "jp", "kr", "pt", "sa"},
		"de": {"ru", "kz", "en", "fr", "cn", "es", "jp", "kr", "pt", "sa"},
		"es": {"ru", "kz", "en", "fr", "cn", "de", "jp", "kr", "pt", "sa"},
		"jp": {"ru", "kz", "en", "fr", "cn", "de", "es", "kr", "pt", "sa"},
		"kr": {"ru", "kz", "en", "fr", "cn", "de", "es", "jp", "pt", "sa"},
		"pt": {"ru", "kz", "en", "fr", "cn", "de", "es", "jp", "kr", "sa"},
		"sa": {"ru", "kz", "en", "fr", "cn", "de", "es", "jp", "kr", "pt"},
	}

	var languages []language.Language
	if err := s.db.Where("code IN ?", allKeys(availableCodes)).Find(&languages).Error; err != nil {
		return err
	}
	langByCode := make(map[string]int, len(languages))
	for _, l := range languages {
		langByCode[l.Code] = l.ID
	}

	var rows []language.AvailableLanguage
	for baseCode, targets := range availableCodes {
		baseID, ok := langByCode[baseCode]
		if !ok {
			continue
		}
		for _, targetCode := range targets {
			targetID, ok := langByCode[targetCode]
			if !ok {
				continue
			}
			rows = append(rows, language.AvailableLanguage{
				BaseLanguageID:   baseID,
				TargetLanguageID: targetID,
			})
		}
	}

	if len(rows) == 0 {
		return nil
	}

	return s.db.CreateInBatches(&rows, 500).Error
}
func (s *DatabaseSeeder) RunUserSeeder() error {

	var ru, kz language.Language

	if err := s.db.Where("code = ?", "ru").First(&ru).Error; err != nil {
		return err
	}
	if err := s.db.Where("code = ?", "kz").First(&kz).Error; err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	email := "drive16052003@gmail.com"
	User := user.User{
		Email:            &email,
		Password:         string(hash),
		Name:             "LingCard",
		BaseLanguageID:   ru.ID,
		TargetLanguageID: kz.ID,
		Role:             role.ROLEADMIN,
		IsBanned:         false,
	}

	return s.db.Create(&User).Error
}
func (s *DatabaseSeeder) RunWordSeeder() error {
	baseFilePath := filepath.Join("", "data", "dataset")

	entries, err := os.ReadDir(baseFilePath)
	if err != nil {
		return err
	}

	var languageDirs []string
	for _, e := range entries {
		if !e.IsDir() || e.Name() == "base" {
			continue
		}
		languageDirs = append(languageDirs, e.Name())
	}

	var langs []language.Language
	if err := s.db.Find(&langs).Error; err != nil {
		return err
	}
	languagesCache := make(map[string]language.Language, len(langs))
	for _, l := range langs {
		languagesCache[l.Code] = l
	}

	for _, languageCode := range languageDirs {
		firstLanguage, ok := languagesCache[languageCode]
		if !ok {
			fmt.Printf("Язык %s не найден в БД\n", languageCode)
			continue
		}

		dirPath := filepath.Join(baseFilePath, languageCode)
		files, err := os.ReadDir(dirPath)
		if err != nil {
			return err
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}
			name := f.Name()
			targetLanguage := strings.TrimSuffix(name, ".jsonl")

			secondLanguage, ok := languagesCache[targetLanguage]
			if !ok {
				fmt.Printf("Язык %s не найден в БД\n", targetLanguage)
				continue
			}

			filePath := filepath.Join(dirPath, name)
			fmt.Printf("Загружаю языковой пакет %s - %s\n", firstLanguage.Code, secondLanguage.Code)

			if err := processFileWithBatchInsert(s.db, filePath, firstLanguage, secondLanguage, languageCode, targetLanguage); err != nil {
				fmt.Printf("Ошибка обработки файла %s: %v\n", filePath, err)
			}
		}
	}

	return nil
}
func (s *DatabaseSeeder) RunPostSeeder() error {
	var usr user.User
	if err := s.db.Where("name = ?", "LingCard").First(&usr).Error; err != nil {
		return err
	}

	codes := []string{"ru", "kz", "en", "es", "fr", "de", "pt", "cn", "jp", "kr", "sa"}
	var langs []language.Language
	if err := s.db.Where("code IN ?", codes).Find(&langs).Error; err != nil {
		return err
	}
	langByCode := make(map[string]int, len(langs))
	for _, l := range langs {
		langByCode[l.Code] = l.ID
	}

	now := time.Now()
	posts := []post.Post{
		{
			Content:       "Всем привет! Рады сообщить, что LingCard открыт для всех желающих изучать языки разных стран. Если интересующего вас языка нет, сообщите нам или проголосуйте за его добавление.",
			Date:          &now,
			Title:         "Стартуем!",
			LanguageID:    langByCode["ru"],
			UserID:        usr.ID,
			Address:       "Россия",
			Status:        post2.APPROVED,
			ViewsCount:    0,
			LikesCount:    0,
			DislikesCount: 0,
		},
		{
			Content:       "Барлығына сәлем! LingCard түрлі елдердің тілдерін үйренгісі келетіндердің бәріне ашық екенін хабарлауға қуаныштымыз. Егер сізді қызықтыратын тіл жоқ болса, бізге хабарлаңыз немесе оны қосуды дауыстап қолдаңыз.",
			Date:          &now,
			Title:         "Бастаймыз!",
			LanguageID:    langByCode["kz"],
			UserID:        usr.ID,
			Address:       "Қазақстан",
			Status:        post2.APPROVED,
			ViewsCount:    0,
			LikesCount:    0,
			DislikesCount: 0,
		},
		{
			Content:       "Hello everyone! We are happy to announce that LingCard is now open to everyone who wants to learn languages from different countries. If your language of interest is not available, let us know or vote for its addition.",
			Date:          &now,
			Title:         "We're starting!",
			LanguageID:    langByCode["en"],
			UserID:        usr.ID,
			Address:       "United Kingdom",
			Status:        post2.APPROVED,
			ViewsCount:    0,
			LikesCount:    0,
			DislikesCount: 0,
		},
		{
			Content:       "¡Hola a todos! Nos complace anunciar que LingCard está abierto para todos aquellos que quieran aprender idiomas de diferentes países. Si tu idioma de interés no está disponible, infórmanos o vota por su inclusión.",
			Date:          &now,
			Title:         "¡Empezamos!",
			LanguageID:    langByCode["es"],
			UserID:        usr.ID,
			Address:       "España",
			Status:        post2.APPROVED,
			ViewsCount:    0,
			LikesCount:    0,
			DislikesCount: 0,
		},
		{
			Content:       "Bonjour à tous! Nous sommes ravis d'annoncer que LingCard est ouvert à tous ceux qui souhaitent apprendre des langues de différents pays. Si votre langue d'intérêt n'est pas disponible, faites-le nous savoir ou votez pour son ajout.",
			Date:          &now,
			Title:         "Nous commençons!",
			LanguageID:    langByCode["fr"],
			UserID:        usr.ID,
			Address:       "France",
			Status:        post2.APPROVED,
			ViewsCount:    0,
			LikesCount:    0,
			DislikesCount: 0,
		},
		{
			Content:       "Hallo zusammen! Wir freuen uns, bekannt zu geben, dass LingCard für alle geöffnet ist, die Sprachen aus verschiedenen Ländern lernen möchten. Wenn Ihre gewünschte Sprache nicht verfügbar ist, teilen Sie uns dies mit oder stimmen Sie für deren Hinzufügung ab.",
			Date:          &now,
			Title:         "Wir starten!",
			LanguageID:    langByCode["de"],
			UserID:        usr.ID,
			Address:       "Deutschland",
			Status:        post2.APPROVED,
			ViewsCount:    0,
			LikesCount:    0,
			DislikesCount: 0,
		},
		{
			Content:       "Olá a todos! Temos o prazer de anunciar que o LingCard está aberto para todos aqueles que desejam aprender idiomas de diferentes países. Se o seu idioma de interesse não estiver disponível, informe-nos ou vote para adicioná-lo.",
			Date:          &now,
			Title:         "Estamos começando!",
			LanguageID:    langByCode["pt"],
			UserID:        usr.ID,
			Address:       "Portugal",
			Status:        post2.APPROVED,
			ViewsCount:    0,
			LikesCount:    0,
			DislikesCount: 0,
		},
		{
			Content:       "大家好！我们很高兴地宣布，LingCard向所有想学习不同国家语言的人开放。如果您感兴趣的语言尚未提供，请告诉我们或投票支持添加。",
			Date:          &now,
			Title:         "我们开始了！",
			LanguageID:    langByCode["cn"],
			UserID:        usr.ID,
			Address:       "中国",
			Status:        post2.APPROVED,
			ViewsCount:    0,
			LikesCount:    0,
			DislikesCount: 0,
		},
		{
			Content:       "皆さん、こんにちは！LingCardが様々な国の言語を学びたいすべての方に開放されたことをお知らせできることを嬉しく思います。ご興味のある言語が利用できない場合は、お知らせいただくか、追加に投票してください。",
			Date:          &now,
			Title:         "スタートします！",
			LanguageID:    langByCode["jp"],
			UserID:        usr.ID,
			Address:       "日本",
			Status:        post2.APPROVED,
			ViewsCount:    0,
			LikesCount:    0,
			DislikesCount: 0,
		},
		{
			Content:       "여러분 안녕하세요! LingCard가 다양한 국가의 언어를 배우고자 하는 모든 분들에게 개방되었음을 알리게 되어 기쁩니다. 관심 있는 언어가 제공되지 않는 경우 저희에게 알려주시거나 추가에 투표해 주세요.",
			Date:          &now,
			Title:         "시작합니다!",
			LanguageID:    langByCode["kr"],
			UserID:        usr.ID,
			Address:       "한국",
			Status:        post2.APPROVED,
			ViewsCount:    0,
			LikesCount:    0,
			DislikesCount: 0,
		},
		{
			Content:       "مرحباً بالجميع! يسعدنا أن نعلن أن LingCard مفتوح للجميع الراغبين في تعلم لغات من مختلف البلدان. إذا كانت لغتك المفضلة غير متوفرة، أخبرنا بذلك أو صوت لإضافتها.",
			Date:          &now,
			Title:         "نحن نبدأ!",
			LanguageID:    langByCode["sa"],
			UserID:        usr.ID,
			Address:       "السعودية",
			Status:        post2.APPROVED,
			ViewsCount:    0,
			LikesCount:    0,
			DislikesCount: 0,
		},
	}

	return s.db.Create(&posts).Error
}
func (s *DatabaseSeeder) RunVoteSeeder() error {
	titleJSON, err := json.Marshal(map[string]string{
		"en": "Select a new language to add!",
		"ru": "Выберите новый язык для добавления!",
		"kz": "Қосу үшін жаңа тілді таңдаңыз!",
		"fr": "Sélectionnez une nouvelle langue à ajouter !",
		"cn": "选择要添加的新语言！",
		"de": "Wählen Sie eine neue Sprache zum Hinzufügen!",
		"es": "¡Selecciona un nuevo idioma para agregar!",
		"jp": "追加する新しい言語を選択してください！",
		"kr": "추가할 새 언어를 선택하세요!",
		"pt": "Selecione um novo idioma para adicionar!",
		"sa": "اختر لغة جديدة لإضافتها!",
	})
	if err != nil {
		return err
	}

	contentJSON, err := json.Marshal(map[string]string{
		"en": "This is the first vote on LingCard! Your opinion is extremely important to us, because it is you who shape the future of our project. We invite you to choose the next language to be added for learning.",
		"ru": "Это первое голосование на LingCard! Ваше мнение крайне важно для нас, потому что именно вы формируете будущее нашего проекта. Мы приглашаем вас выбрать следующий язык для изучения.",
		"kz": "Бұл LingCard-тағы алғашқы дауыс беру! Сіздің пікіріңіз біз үшін өте маңызды, өйткені дәл сіз біздің жобамыздың болашағын қалыптастырасыз. Сізді оқу үшін келесі тілді таңдауға шақырамыз.",
		"fr": "C'est le premier vote sur LingCard ! Votre avis est extrêmement important pour nous, car c'est vous qui façonnez l'avenir de notre projet. Nous vous invitons à choisir la prochaine langue à ajouter pour l'apprentissage.",
		"cn": "这是LingCard上的第一次投票！您的意见对我们极为重要，因为正是您塑造了我们项目的未来。我们邀请您选择下一个要添加学习的语言。",
		"de": "Dies ist die erste Abstimmung auf LingCard! Ihre Meinung ist uns äußerst wichtig, denn Sie sind es, die die Zukunft unseres Projekts gestalten. Wir laden Sie ein, die nächste Sprache zum Lernen auszuwählen.",
		"es": "¡Esta es la primera votación en LingCard! Tu opinión es extremadamente importante para nosotros, porque eres tú quien da forma al futuro de nuestro proyecto. Te invitamos a elegir el próximo idioma para aprender.",
		"jp": "これはLingCardでの最初の投票です！あなたのご意見は私たちにとって非常に重要です。なぜなら、あなたが私たちのプロジェクトの未来を形作るからです。学習用に追加する次の言語を選択してください。",
		"kr": "이것은 LingCard의 첫 번째 투표입니다! 귀하의 의견은 우리에게 매우 중요합니다. 귀하가 우리 프로젝트의 미래를 만들기 때문입니다. 학습에 추가할 다음 언어를 선택해 주세요.",
		"pt": "Esta é a primeira votação no LingCard! Sua opinião é extremamente importante para nós, porque é você quem molda o futuro do nosso projeto. Convidamos você a escolher o próximo idioma a ser adicionado para aprendizado.",
		"sa": "هذا هو التصويت الأول على LingCard! رأيك مهم جداً بالنسبة لنا، لأنك أنت من تشكل مستقبل مشروعنا. ندعوك لاختيار اللغة التالية لإضافتها للتعلم.",
	})
	if err != nil {
		return err
	}

	vote := voice.Vote{
		Title:    string(titleJSON),
		Content:  string(contentJSON),
		IsActive: true,
	}
	if err := s.db.Create(&vote).Error; err != nil {
		return err
	}

	var languages []language.Language
	if err := s.db.Where("is_active = ?", false).Find(&languages).Error; err != nil {
		return err
	}

	if len(languages) == 0 {
		return nil
	}

	options := make([]voice.VoteOption, 0, len(languages))
	for _, lang := range languages {
		optionContent, err := json.Marshal(map[string]string{
			"picture": "/flags/" + lang.Code + ".svg",
			"code":    lang.Code,
		})
		if err != nil {
			return err
		}
		options = append(options, voice.VoteOption{
			VoteID:  vote.ID,
			Title:   lang.Name,
			Content: string(optionContent),
		})
	}

	return s.db.Create(&options).Error
}

func allKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
func processFileWithBatchInsert(
	db *gorm.DB,
	filePath string,
	firstLanguage language.Language,
	secondLanguage language.Language,
	sourceField string,
	targetField string,
) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл %s: %w", filePath, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1024*1024), 16*1024*1024)

	var wordsBatch []word.Word
	var translationsBatch []word.WordTranslation
	lineNumber := 0

	flush := func() error {
		if len(wordsBatch) == 0 {
			return nil
		}
		if err := insertBatch(db, wordsBatch, translationsBatch); err != nil {
			return err
		}
		wordsBatch = wordsBatch[:0]
		translationsBatch = translationsBatch[:0]
		return nil
	}

	for scanner.Scan() {
		lineNumber++
		line := scanner.Bytes()

		var wrd map[string]any
		if err := json.Unmarshal(line, &wrd); err != nil {
			fmt.Printf("Ошибка в строке %d: %v\n", lineNumber, err)
			continue
		}

		src, ok := wrd[sourceField]
		if !ok || src == nil {
			fmt.Printf("Ошибка в строке %d: нет поля %s\n", lineNumber, sourceField)
			continue
		}
		srcStr, _ := src.(string)

		// Transcription — string, поэтому приводим безопасно
		transcription := ""
		if t, ok := wrd["transcription"].(string); ok {
			transcription = t
		}

		// Level — int, поэтому парсим из float64 (json.Number не используется)
		level := 0
		if l, ok := wrd["level"].(float64); ok {
			level = int(l)
		}

		wordsBatch = append(wordsBatch, word.Word{
			Text:          srcStr,
			Transcription: transcription,
			LanguageID:    firstLanguage.ID,
			Level:         level,
		})

		translation := ""
		if t, ok := wrd[targetField].(string); ok {
			translation = t
		}

		translationsBatch = append(translationsBatch, word.WordTranslation{
			Translation:      translation,
			TargetLanguageID: secondLanguage.ID,
		})

		if len(wordsBatch) >= batchSize {
			if err := flush(); err != nil {
				fmt.Printf("Ошибка при пакетной вставке: %v\n", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return flush()
}

func insertBatch(db *gorm.DB, words []word.Word, translations []word.WordTranslation) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&words).Error; err != nil {
			return err
		}

		rows := make([]word.WordTranslation, 0, len(translations))
		for i := range translations {
			if i >= len(words) || words[i].ID == 0 {
				continue
			}
			translations[i].WordID = words[i].ID
			rows = append(rows, translations[i])
		}

		if len(rows) == 0 {
			return nil
		}

		return tx.Create(&rows).Error
	})
}
