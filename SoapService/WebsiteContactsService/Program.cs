using SoapCore;
using SoapCore.Extensibility;
using WebsiteContactsService.clases;
using WebsiteContactsService.Contracts;
using WebsiteContactsService.Models;
using WebsiteContactsService.Services;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddAuthorization();
builder.Services.AddHttpClient();

builder.Services.AddSingleton<RapidApiService>();
builder.Services.AddSingleton<IContactSearchService, ContactSearchService>();
builder.Services.AddSingleton<IFaultExceptionTransformer, FaultExceptionTransformer>();

var app = builder.Build();

app.UseAuthorization();
app.UseSoapEndpoint<IContactSearchService>("/contact", new SoapEncoderOptions());


app.Run();