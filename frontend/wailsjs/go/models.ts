export namespace main {
	
	export class Attachment {
	    name: string;
	    size: number;
	    formattedSize: string;
	
	    static createFrom(source: any = {}) {
	        return new Attachment(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.size = source["size"];
	        this.formattedSize = source["formattedSize"];
	    }
	}
	export class Email {
	    seqNum: number;
	    from: string;
	    subject: string;
	    date: string;
	    isRead: boolean;
	    isStarred: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Email(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.seqNum = source["seqNum"];
	        this.from = source["from"];
	        this.subject = source["subject"];
	        this.date = source["date"];
	        this.isRead = source["isRead"];
	        this.isStarred = source["isStarred"];
	    }
	}
	export class FullEmail {
	    seqNum: number;
	    from: string;
	    to: string;
	    cc: string;
	    subject: string;
	    body: string;
	    attachments: Attachment[];
	
	    static createFrom(source: any = {}) {
	        return new FullEmail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.seqNum = source["seqNum"];
	        this.from = source["from"];
	        this.to = source["to"];
	        this.cc = source["cc"];
	        this.subject = source["subject"];
	        this.body = source["body"];
	        this.attachments = this.convertValues(source["attachments"], Attachment);
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

