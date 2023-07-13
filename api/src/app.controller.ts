import { Controller, Get, Headers } from '@nestjs/common';
import { UserDataClient } from './data/user.data';

@Controller('users')
export class AppController {
  constructor(private readonly userDataClient: UserDataClient) {}

  @Get()
  async get(@Headers('USER-ID') id?: string): Promise<string> {
    const user = await this.userDataClient.getById(id);
    /*    let result = '';
    users.forEach((element) => {
      result += element.email;
    });
    return result;
  }
*/
    return user.email + 'aaa';
  }
}
