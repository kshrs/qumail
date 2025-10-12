export namespace main {
	
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
	    from: string;
	    to: string;
	    cc: string;
	    subject: string;
	    body: string;
	
	    static createFrom(source: any = {}) {
	        return new FullEmail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.from = source["from"];
	        this.to = source["to"];
	        this.cc = source["cc"];
	        this.subject = source["subject"];
	        this.body = source["body"];
	    }
	}

}

