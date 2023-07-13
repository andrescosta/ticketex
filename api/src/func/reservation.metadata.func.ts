import { Injectable } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { RpcClient } from 'src/common/rpc.client';
import { ReservationMetadata } from 'src/models/reservation.metadata';

@Injectable()
export class ReservationMetadataFuncClient extends RpcClient<ReservationMetadata> {
  constructor(httpService: HttpService) {
    super(httpService, 'http://localhost:8080/v1/reservations/metadata');
  }
}
