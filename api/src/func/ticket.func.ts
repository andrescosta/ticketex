import { Injectable } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { RpcClient } from 'src/common/rpc.client';
import { Ticket } from 'src/models/ticket';

@Injectable()
export class TicketFuncClient extends RpcClient<Ticket> {
  constructor(httpService: HttpService) {
    super(httpService, 'http://localhost:8082/v1/tickets/');
  }
}
