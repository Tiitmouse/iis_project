export namespace api {
	
	export class ArrayOfstring {
	    string?: string[];
	
	    static createFrom(source: any = {}) {
	        return new ArrayOfstring(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.string = source["string"];
	    }
	}
	export class ContactRecord {
	    RecordType?: string;
	    Sources?: ArrayOfstring;
	    Value?: string;
	
	    static createFrom(source: any = {}) {
	        return new ContactRecord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.RecordType = source["RecordType"];
	        this.Sources = this.convertValues(source["Sources"], ArrayOfstring);
	        this.Value = source["Value"];
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
	export class ArrayOfContactRecord {
	    ContactRecord?: ContactRecord[];
	
	    static createFrom(source: any = {}) {
	        return new ArrayOfContactRecord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ContactRecord = this.convertValues(source["ContactRecord"], ContactRecord);
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

