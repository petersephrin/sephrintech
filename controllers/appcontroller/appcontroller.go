package appcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	mail "github.com/xhit/go-simple-mail/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"title": "Sephrin Tech",
	}, "layouts/master")
}

func About(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"title": "About Sephrin Tech",
	}, "layouts/master")
}

func Portfolio(c *fiber.Ctx) error {
	return c.Render("portfolio", fiber.Map{
		"title": "Talk to Us-Sephrin Tech",
	}, "layouts/master")
}

func Contact(c *fiber.Ctx) error {
	return c.Render("contact", fiber.Map{
		"title": "Talk to Us-Sephrin Tech",
	}, "layouts/master")
}

func SendEmailContact(c *fiber.Ctx) error {
	name := c.FormValue("inputName")
	emailOfCustomer := c.FormValue("inputEmail")
	phoneNumber := c.FormValue("inputPhone")
	projectType := c.FormValue("projectType")
	projectDescription := c.FormValue("projectDescription")

	logrus.Info(name + emailOfCustomer + phoneNumber + projectType + projectDescription)

	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = "sephrintech@gmail.com"
	server.Password = "bueolmbogxiadtpn"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		logrus.Warn(err)
		c.Redirect("/contact")
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("Sephrin Tech <sephrintech@gmail.com>")
	email.AddTo("safaripeterw@gmail.com")
	//email.AddCc("another_you@example.com")
	email.SetSubject("New Client(" + name + ") For SephrinTech")

	email.SetBody(mail.TextPlain, "Hi Sir, A new client just reached out. They want a "+projectType+". Their email is: "+emailOfCustomer+" and their phone is :"+phoneNumber+". Their project description is "+projectDescription)
	//email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		logrus.Warn(err)
		c.RedirectBack("/contact")
	}

	return nil
}
