export namespace api {
	
	export class CityWeatherInfo {
	    City: string;
	    Temperature: number;
	    WeatherCondition: string;
	
	    static createFrom(source: any = {}) {
	        return new CityWeatherInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.City = source["City"];
	        this.Temperature = source["Temperature"];
	        this.WeatherCondition = source["WeatherCondition"];
	    }
	}
	export class SoapPhoneEntry {
	    XMLName: xml.Name;
	    Value: string;
	    Sources: string[];
	
	    static createFrom(source: any = {}) {
	        return new SoapPhoneEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.XMLName = this.convertValues(source["XMLName"], xml.Name);
	        this.Value = source["Value"];
	        this.Sources = source["Sources"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class SoapEmailEntry {
	    XMLName: xml.Name;
	    Value: string;
	    Sources: string[];
	
	    static createFrom(source: any = {}) {
	        return new SoapEmailEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.XMLName = this.convertValues(source["XMLName"], xml.Name);
	        this.Value = source["Value"];
	        this.Sources = source["Sources"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class SoapContactRecord {
	    XMLName: xml.Name;
	    Domain: string;
	    Query: string;
	    Emails: SoapEmailEntry[];
	    PhoneNumbers: SoapPhoneEntry[];
	    Facebook: string;
	    Instagram: string;
	    Github: string;
	    Linkedin: string;
	    Twitter: string;
	    Youtube: string;
	    Pinterest: string;
	    Tiktok: string;
	    Snapchat: string;
	
	    static createFrom(source: any = {}) {
	        return new SoapContactRecord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.XMLName = this.convertValues(source["XMLName"], xml.Name);
	        this.Domain = source["Domain"];
	        this.Query = source["Query"];
	        this.Emails = this.convertValues(source["Emails"], SoapEmailEntry);
	        this.PhoneNumbers = this.convertValues(source["PhoneNumbers"], SoapPhoneEntry);
	        this.Facebook = source["Facebook"];
	        this.Instagram = source["Instagram"];
	        this.Github = source["Github"];
	        this.Linkedin = source["Linkedin"];
	        this.Twitter = source["Twitter"];
	        this.Youtube = source["Youtube"];
	        this.Pinterest = source["Pinterest"];
	        this.Tiktok = source["Tiktok"];
	        this.Snapchat = source["Snapchat"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

export namespace main {
	
	export class JsFile {
	    Name: string;
	    Size: number;
	    Type: string;
	
	    static createFrom(source: any = {}) {
	        return new JsFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Size = source["Size"];
	        this.Type = source["Type"];
	    }
	}
	export class Response {
	    Data: any;
	    Error: any;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Data = source["Data"];
	        this.Error = source["Error"];
	    }
	}

}

export namespace xml {
	
	export class Name {
	    Space: string;
	    Local: string;
	
	    static createFrom(source: any = {}) {
	        return new Name(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Space = source["Space"];
	        this.Local = source["Local"];
	    }
	}

}

