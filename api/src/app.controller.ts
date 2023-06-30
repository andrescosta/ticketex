import { Controller, Get } from '@nestjs/common';
import { DataService } from './data.service';
import { FuncService } from './func.service';

@Controller('users')
export class AppController {
  constructor(
    private readonly dataService: DataService,
    private readonly funcService: FuncService,
  ) {}

  @Get()
  getHello(): string {
    return this.dataService.getHello() + this.funcService.getHello();
  }
}
