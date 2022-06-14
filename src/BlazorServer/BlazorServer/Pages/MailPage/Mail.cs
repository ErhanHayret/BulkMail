using BlazorServer.Data.Models;
using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Components.Server.ProtectedBrowserStorage;

namespace BlazorServer.Pages.MailPage
{
    public partial class Mail
    {
        public string TitleString
        {
            get
            {
                getStorage();
                return "Toplu Mail";
            }
        }
        public string SenderEmail { get; set; }
        public string Password { get; set; }
        public string Subject { get; set; }
        public string Message { get; set; }
        public string MailException { get; set; }
        [Inject]
        private ProtectedLocalStorage localStore { get; set; }
        private string userName;

        public void Send()
        {
            if (!string.IsNullOrEmpty(SenderEmail) && !string.IsNullOrEmpty(Password) && !string.IsNullOrEmpty(Subject) && !string.IsNullOrEmpty(Message) && ArriveList.Count > 0)
            {
                string mailText = "Content-Type: text/plain; charset=utf-8\nFrom:" + userName + " <" + SenderEmail + ">\nTo: YOU\nSubject: " + Subject + "\n\n" + Message;
                MailModel mail = new MailModel() { MailText = mailText, SenderEmail = this.SenderEmail, SenderEmailPsw = this.Password, ArriveEmails = ArriveList };
                using (var httpClient = new HttpClient())
                {
                    var response = httpClient.PostAsJsonAsync<MailModel>("http://localhost:10000/Mail/SendMail", mail).Result;
                    if (response.IsSuccessStatusCode)
                    {
                        if (response.StatusCode == System.Net.HttpStatusCode.OK)
                        {
                            MailException = "Gönderim Başarılı";
                        }
                        else
                        {
                            MailException = response.StatusCode.ToString();
                        }
                    }
                    else
                    {
                        MailException = "Bir hata oluştu tekrar deneyiniz";
                    }
                }
            }
            else
            {
                MailException = "Eksik bilgileri doldurunuz";
            }
        }

        private async void getStorage()
        {
            var result = await localStore.GetAsync<string>("UserName");
            userName = result.Success ? result.Value : "";
            if (string.IsNullOrEmpty(userName))
            {
                navManager.NavigateTo("login/login");
            }
        }
    }
}
