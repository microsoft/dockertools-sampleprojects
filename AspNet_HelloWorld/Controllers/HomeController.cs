using System;
using Microsoft.AspNet.Mvc;

namespace MvcSample.Web
{
    public class HomeController : Controller
    {
        public IActionResult Index()
        {
            ViewBag.OS =  "OS: " + Environment.GetEnvironmentVariable("OS");
            return View();
        }
    }
}
