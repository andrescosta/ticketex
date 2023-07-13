import { Injectable } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { RpcClient } from 'src/common/rpc.client';
import { Message } from 'src/models/message';

@Injectable()
export class MessagingFuncClient extends RpcClient<Message> {
  constructor(httpService: HttpService) {
    super(httpService, 'http://localhost:8585/v1/messaging/');
  }
}
