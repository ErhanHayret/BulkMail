using BlazorServer.Data.Models;
using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Components.Server.ProtectedBrowserStorage;

namespace BlazorServer.Pages.LoginPage
{
    public partial class Login
    {
        public string UserName { get; set; }
        public string Password { get; set; }
        public string UserException { get; set; }

        [Inject]
        private ProtectedLocalStorage localStore { get; set; }
        private UserModel user = new UserModel();

        public async void LoginClick()
        {
            user.UserName = this.UserName;
            user.Password = this.Password;
            using (HttpClient client = new HttpClient())
            {
                var response = client.PostAsJsonAsync<UserModel>("http://localhost:10000/User/GetUser", user).Result;
                if (response.IsSuccessStatusCode)
                {
                    user = response.Content.ReadFromJsonAsync<UserModel>().Result;
                    if (!string.IsNullOrEmpty(user.UserName))
                    {
                        await localStore.SetAsync("UserName", user.UserName);
                        navManager.NavigateTo("mail/mail");
                    }
                    else
                    {
                        UserException = "Kullanıcı veya Şifre Hatalı";
                    }
                }
                else
                {
                    UserException = "Bir sorun oluştu lütfen tekrar deneyiniz";
                }
            }

        }
    }
}
