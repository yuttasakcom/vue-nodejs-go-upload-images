import * as Koa from 'koa'
import * as Router from '@koa/router'
import * as cors from '@koa/cors'
import * as multer from '@koa/multer'
import { bodyParser } from '@koa/bodyparser'

const app = new Koa()
const router = new Router()
const upload = multer()

app.use(cors())
app.use(bodyParser())

// router.post('/upload', async (ctx) => {
//   console.log(ctx.request)
//   ctx.body = 'uploaded'
// })

router.post(
  '/upload',
  upload.fields([
    {
      name: 'id_card_image',
      maxCount: 1
    }
  ]),
  (ctx) => {
    console.log('id_card_image', ctx.files.id_card_image[0])
    console.log('id_card_number', ctx.request.body.id_card_number)
    ctx.body = 'done'
  }
)

app.use(router.routes())

app.listen(3000)

console.log('Server running on port 3000')
