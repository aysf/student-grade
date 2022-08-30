package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Niko",
			LastName:  "Budi",
			Grades: []Grade{
				{
					Title: "PR Fisika",
					Type:  HomeWork,
					Score: 9.4,
				},
				{
					Title: "Quiz Fisia",
					Type:  Quiz,
					Score: 9.4,
				},
				{
					Title: "Ujian Fisika",
					Type:  Test,
					Score: 9.4,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Rico",
			LastName:  "Siaga",
			Grades: []Grade{
				{
					Title: "PR Matematika",
					Type:  HomeWork,
					Score: 9.4,
				},
				{
					Title: "Quiz IPA",
					Type:  Quiz,
					Score: 2.4,
				},
				{
					Title: "Ujian Fisika",
					Type:  Test,
					Score: 9.9,
				},
			},
		},
		{
			ID:        3,
			FirstName: "Kaka",
			LastName:  "Mesi",
			Grades: []Grade{
				{
					Title: "PR Fisika",
					Type:  HomeWork,
					Score: 8.4,
				},
				{
					Title: "Quiz Fisia",
					Type:  Quiz,
					Score: 6.4,
				},
				{
					Title: "Ujian Kimia",
					Type:  Test,
					Score: 5.4,
				},
			},
		},
		{
			ID:        4,
			FirstName: "Ruben",
			LastName:  "Gerald",
			Grades: []Grade{
				{
					Title: "PR Bio",
					Type:  HomeWork,
					Score: 5.4,
				},
				{
					Title: "Quiz Ekonomi",
					Type:  Quiz,
					Score: 6.4,
				},
				{
					Title: "Ujian Geo",
					Type:  Test,
					Score: 7.4,
				},
			},
		},
		{
			ID:        5,
			FirstName: "Jack",
			LastName:  "Ko",
			Grades: []Grade{
				{
					Title: "PR Matematika",
					Type:  HomeWork,
					Score: 5.4,
				},
				{
					Title: "Quiz Fisia",
					Type:  Quiz,
					Score: 6.4,
				},
				{
					Title: "Ujian Ekonomi",
					Type:  Test,
					Score: 8.1,
				},
			},
		},
		{
			ID:        6,
			FirstName: "Ani",
			LastName:  "Fadilah",
			Grades: []Grade{
				{
					Title: "PR Jerman",
					Type:  HomeWork,
					Score: 9.4,
				},
				{
					Title: "Quiz China",
					Type:  Quiz,
					Score: 9.4,
				},
				{
					Title: "Ujian Jepang",
					Type:  Test,
					Score: 9.4,
				},
			},
		},
		{
			ID:        7,
			FirstName: "Komang",
			LastName:  "Made",
			Grades: []Grade{
				{
					Title: "PR Art",
					Type:  HomeWork,
					Score: 9.9,
				},
				{
					Title: "Quiz Fisia",
					Type:  Quiz,
					Score: 3.4,
				},
				{
					Title: "Ujian Matematika",
					Type:  Test,
					Score: 5.4,
				},
			},
		},
		{
			ID:        8,
			FirstName: "John",
			LastName:  "Pillow",
			Grades: []Grade{
				{
					Title: "PR Fisika",
					Type:  HomeWork,
					Score: 6.3,
				},
				{
					Title: "Quiz Fisia",
					Type:  Quiz,
					Score: 5.8,
				},
				{
					Title: "Ujian Matematika",
					Type:  Test,
					Score: 7.2,
				},
			},
		},
		{
			ID:        9,
			FirstName: "Ken",
			LastName:  "Maru",
			Grades: []Grade{
				{
					Title: "PR Fisika",
					Type:  HomeWork,
					Score: 9.4,
				},
				{
					Title: "Quiz Fisia",
					Type:  Quiz,
					Score: 9.4,
				},
				{
					Title: "Ujian Fisika",
					Type:  Test,
					Score: 3.4,
				},
			},
		},
		{
			ID:        10,
			FirstName: "Ben",
			LastName:  "So",
			Grades: []Grade{
				{
					Title: "PR Kimia",
					Type:  HomeWork,
					Score: 5.4,
				},
				{
					Title: "Quiz Fisia",
					Type:  Quiz,
					Score: 9.4,
				},
				{
					Title: "Ujian Fisika",
					Type:  Test,
					Score: 9.4,
				},
			},
		},
	}
}
