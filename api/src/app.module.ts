import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { DataService } from './data/data.service';
import { FuncService } from './func/func.service';

@Module({
  imports: [],
  controllers: [AppController],
  providers: [DataService, FuncService],
})
export class AppModule {}
