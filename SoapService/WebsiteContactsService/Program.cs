using SoapCore;
using WebsiteContactsService.Contracts;
using WebsiteContactsService.Services;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddAuthorization();
builder.Services.AddHttpClient();

builder.Services.AddSingleton<RapidApiService>();
builder.Services.AddSingleton<IContactSearchService, ContactSearchService>();

var app = builder.Build();

app.UseAuthorization();
app.UseSoapEndpoint<IContactSearchService>("/contact", new SoapEncoderOptions());


app.Run();