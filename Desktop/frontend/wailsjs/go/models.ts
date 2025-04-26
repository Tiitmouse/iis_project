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

