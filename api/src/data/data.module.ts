import { HttpModule } from '@nestjs/axios';
import { Module } from '@nestjs/common/decorators';
import { UsersService } from './users.service';
import { PartnersService } from './partners.service';
import { AdventuresService } from './adventures.service';

@Module({
  imports: [
    HttpModule.registerAsync({
      useFactory: () => ({
        timeout: 5000,
        maxRedirects: 5,
      }),
    }),
  ],
  providers: [UsersService, PartnersService, AdventuresService],
})
export class DataModule {}
