import { HttpModule } from '@nestjs/axios';
import { Module } from '@nestjs/common/decorators';
import { ReservationMetadataFuncClient } from './reservation.metadata.func';
import { MessagingFuncClient } from './messaging.func';
import { ReservationFuncClient } from './reservation.func';
import { TicketFuncClient } from './ticket.func';

@Module({
  imports: [
    HttpModule.registerAsync({
      useFactory: () => ({
        timeout: 5000,
        maxRedirects: 5,
      }),
    }),
  ],
  providers: [
    ReservationMetadataFuncClient,
    MessagingFuncClient,
    ReservationFuncClient,
    TicketFuncClient,
  ],
  exports: [
    ReservationMetadataFuncClient,
    MessagingFuncClient,
    ReservationFuncClient,
    TicketFuncClient,
  ],
})
export class FuncModule {}
