using BlazorServer.Data.Models;
using System.Net;

namespace BlazorServer.Pages.MailsPage
{
    public partial class Mails
    {
        public Mails()
        {
            using (var client = new HttpClient())
            {
                var response = client.GetAsync("http://localhost:10000/Mail/GetAllMails").Result;
                if (response.StatusCode == HttpStatusCode.OK)
                {
                    MailList=response.Content.ReadFromJsonAsync<List<MailModel>>().Result;
                }
            }
        }

        List<MailModel> MailList;
    }
}
