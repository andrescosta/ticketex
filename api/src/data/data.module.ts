import { HttpModule } from '@nestjs/axios';
import { Module } from '@nestjs/common/decorators';
import { UserDataClient } from './user.data';
import { PartnerDataClient } from './partner.data';
import { AdventureDataClient } from './adventure.data';

@Module({
  imports: [
    HttpModule.registerAsync({
      useFactory: () => ({
        timeout: 5000,
        maxRedirects: 5,
      }),
    }),
  ],
  providers: [UserDataClient, PartnerDataClient, AdventureDataClient],
  exports: [UserDataClient, PartnerDataClient, AdventureDataClient],
})
export class DataModule {}
