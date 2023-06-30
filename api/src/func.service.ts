import { Injectable } from '@nestjs/common';

@Injectable()
export class FuncService {
  getHello(): string {
    return 'Hello World from func!';
  }
}
