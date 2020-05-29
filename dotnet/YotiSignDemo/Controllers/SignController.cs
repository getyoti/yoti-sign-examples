using System;
using System.IO;
using System.Net.Http;
using System.Text;
using System.Text.Encodings.Web;
using System.Text.Json;
using Microsoft.AspNetCore.Mvc;

namespace YotiSignDemo.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class SignController : ControllerBase
    {
        [HttpGet]
        public IActionResult Get()
        {
            Uri apiUri = GetRequestUrl();

            string yotiAuthenticationToken = Environment.GetEnvironmentVariable("YOTI_AUTHENTICATION_TOKEN");

            if (string.IsNullOrEmpty(yotiAuthenticationToken))
            {
                return Unauthorized("YOTI_AUTHENTICATION_TOKEN not found");
            }

            string optionsJson;
            using (StreamReader r = System.IO.File.OpenText("options.json"))
            {
                optionsJson = r.ReadToEnd();
            }

            using (var pdfFileContent = new StreamContent(
                new FileStream(
                    "test.pdf",
                    FileMode.Open,
                    FileAccess.Read)))
            {
                using (var optionsJsonContent = new StringContent(
                    optionsJson,
                    Encoding.UTF8))
                {
                    MultipartFormDataContent multiPartContent = new MultipartFormDataContent
                    {
                        { pdfFileContent, "file", "test.pdf" },
                        { optionsJsonContent, "options" }
                    };

                    using (HttpRequestMessage request = new HttpRequestMessage
                    {
                        Method = HttpMethod.Post,
                        RequestUri = apiUri,
                        Content = multiPartContent
                    })
                    {
                        request.Headers.Add("Authorization", "Bearer " + yotiAuthenticationToken);

                        using (HttpClient httpClient = new HttpClient())
                        {
                            HttpResponseMessage response = httpClient.SendAsync(request).Result;

                            string responseContent = response.Content.ReadAsStringAsync().Result;

                            return new JsonResult(
                                responseContent,
                                new JsonSerializerOptions
                                {
                                    WriteIndented = true,
                                    Encoder = JavaScriptEncoder.UnsafeRelaxedJsonEscaping
                                });
                        }
                    }
                }
            }
        }

        private Uri GetRequestUrl()
        {
            if (!string.IsNullOrEmpty(Environment.GetEnvironmentVariable("YOTI_SIGN_BASE_URL")))
            {
                return new Uri(Environment.GetEnvironmentVariable("YOTI_SIGN_BASE_URL"));
            }
            else
            {
                return new Uri("https://api.yotisign.com/v2/envelopes");
            }
        }
    }
}