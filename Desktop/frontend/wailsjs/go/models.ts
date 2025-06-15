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
	export class SoapContactRecord {
	    RecordType: string;
	    Value: string;
	    Sources: string[];
	
	    static createFrom(source: any = {}) {
	        return new SoapContactRecord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.RecordType = source["RecordType"];
	        this.Value = source["Value"];
	        this.Sources = source["Sources"];
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

