import * as Koa from 'koa'
import * as Router from '@koa/router'
import * as cors from '@koa/cors'
import * as multer from '@koa/multer'
import * as FormData from 'form-data'
import { bodyParser } from '@koa/bodyparser'
import axios from 'axios'

const app = new Koa()
const router = new Router()
const upload = multer()

app.use(cors())
app.use(bodyParser())

router.post(
  '/upload',
  upload.fields([
    {
      name: 'id_card_image',
      maxCount: 1
    },
    {
      name: 'bank_image',
      maxCount: 1
    }
  ]),
  async (ctx) => {
    const form = new FormData()
    const idCardNumber = ctx.request.body.id_card_number
    const idCardImage = ctx.files.id_card_image[0]
    const bankNumber = ctx.request.body.bank_number
    const bankImage = ctx.files.bank_image[0]

    form.append('id_card_number', idCardNumber)
    form.append('id_card_image', idCardImage.buffer, {
      filename: idCardImage.originalname,
      contentType: idCardImage.mimetype
    })
    form.append('bank_number', bankNumber)
    form.append('bank_image', bankImage.buffer, {
      filename: bankImage.originalname,
      contentType: bankImage.mimetype
    })

    const res = await axios.post('http://0.0.0.0:8081/upload', form, {
      headers: form.getHeaders()
    })

    ctx.body = res.data
  }
)

app.use(router.routes())

app.listen(3000)

console.log('Server running on port 3000')
