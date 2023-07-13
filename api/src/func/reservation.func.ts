import { Injectable } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { RpcClient } from 'src/common/rpc.client';
import { Reservation } from 'src/models/reservation';

@Injectable()
export class ReservationFuncClient extends RpcClient<Reservation> {
  constructor(httpService: HttpService) {
    super(httpService, 'http://localhost:8080/v1/reservations/');
  }
}
