export namespace main {
	
	export class Email {
	    from: string;
	    subject: string;
	    date: string;
	
	    static createFrom(source: any = {}) {
	        return new Email(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.from = source["from"];
	        this.subject = source["subject"];
	        this.date = source["date"];
	    }
	}

}

