export namespace main {
	
	export class Email {
	    seqNum: number;
	    from: string;
	    subject: string;
	    date: string;
	    isRead: boolean;
	
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
	    }
	}

}

