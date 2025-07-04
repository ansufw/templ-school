package handlers

import (
	"net/http"

	"templ-school/web/entity"
	"templ-school/web/templates/pages"
)

// menuData holds the initialized data structure, mirroring the provided JSON.
var menuData = []entity.MenuCategory{
	{
		Title: "MENU",
		Items: []entity.MenuItem{
			{Icon: "/home.png", Label: "Home", Href: "/"},
			{Icon: "/teacher.png", Label: "Teachers", Href: "/list/teachers"},
			{Icon: "/student.png", Label: "Students", Href: "/list/students"},
			{Icon: "/parent.png", Label: "Parents", Href: "/list/parents"},
			{Icon: "/subject.png", Label: "Subjects", Href: "/list/subjects"},
			{Icon: "/class.png", Label: "Classes", Href: "/list/classes"},
			{Icon: "/lesson.png", Label: "Lessons", Href: "/list/lessons"},
			{Icon: "/exam.png", Label: "Exams", Href: "/list/exams"},
			{Icon: "/assignment.png", Label: "Assignments", Href: "/list/assignments"},
			{Icon: "/result.png", Label: "Results", Href: "/list/results"},
			{Icon: "/attendance.png", Label: "Attendance", Href: "/list/attendance"},
			{Icon: "/calendar.png", Label: "Events", Href: "/list/events"},
			{Icon: "/message.png", Label: "Messages", Href: "/list/messages"},
			{Icon: "/announcement.png", Label: "Announcements", Href: "/list/announcements"},
		},
	},
	{
		Title: "OTHER",
		Items: []entity.MenuItem{
			{Icon: "/profile.png", Label: "Profile", Href: "/profile"},
			{Icon: "/setting.png", Label: "Settings", Href: "/settings"},
			{Icon: "/logout.png", Label: "Logout", Href: "/logout"},
		},
	},
}

func AdminHome(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	component := pages.Admin(menuData)
	component.Render(r.Context(), w)
}
