import 'dotenv/config'
import * as Amqp from "amqp-ts"
import {injectable} from "inversify"
import {Logger} from "tslog"

const log: Logger<any> = new Logger()

@injectable()
class AMQPController {
  public connection: Amqp.Connection

  constructor() {
    const MQ_RABBIT_URI = process.env.MQ_RABBIT_URI

    this.connection = new Amqp.Connection(MQ_RABBIT_URI)

    this.connection.on('open_connection', () => log.info('open_connection'))
    this.connection.on('trying_connect', () => log.warn('trying_connect'))
    this.connection.on('error_connection', () => log.warn('error_connection'))
    this.connection.on('close_connection', () => log.warn('close_connection'))
  }
}

export default AMQPController
