namespace BlazorServer.Data.Models
{
    public class MailModel
    {
        public string Id { get; set; }

        public string MailText { get; set; }

        public string SenderEmail { get; set; }

        public string SenderEmailPsw { get; set; }

        public List<string> ArriveEmails { get; set; }
    }
}
