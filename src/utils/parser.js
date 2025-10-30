// parser.js
import { Console } from 'console';

const console = new Console(process.stdout);

interface Request {
  path: string;
  method: string;
  body: any;
}

interface Response {
  status: number;
  body: any;
}

class Parser {
  private responses: Response[] = [];

  public addResponse(status: number, body: any) {
    this.responses.push({ status, body });
  }

  public getResponses(): Response[] {
    return this.responses;
  }

  public clearResponses() {
    this.responses = [];
  }

  public getResponse(index: number): Response {
    if (index < this.responses.length) {
      return this.responses[index];
    } else {
      return null;
    }
  }
}

class RequestHandler {
  private parser: Parser = new Parser();
  private request: Request | null = null;

  public handleRequest(req: Request) {
    this.request = req;
    try {
      const response: Response = this.processRequest();
      this.parser.addResponse(response.status, response.body);
    } catch (error) {
      console.error(error);
    }
  }

  public getResponse(index: number): Response | null {
    return this.parser.getResponse(index);
  }

  public clearResponses() {
    this.parser.clearResponses();
  }

  private processRequest(): Response {
    // Logic for processing the request goes here
    // For demonstration purposes, let's return a simple response
    return { status: 200, body: 'Response from server' };
  }
}

const handler = new RequestHandler();
const request: Request = {
  path: '/path',
  method: 'GET',
  body: {},
};

handler.handleRequest(request);
console.log(JSON.stringify(handler.getResponse(0)));