package controller

const CONTROLLER_HEADER = `
package {{Package}}

import (
    "context"

    "{{Root}}/models"
)`
const CONTROLLER_INDEX = `func (u *{{Controller}}) Index(c context.Context, page int) ([]models.{{Model}}, int, error) {
    items, total, err := u.db.ListAll(page)
    if err != nil {
        return nil, 0, err
    }

    return items, total, nil
}`

const CONTROLLER_CREATE = `func (u *{{Controller}}) Store(c context.Context, r *StoreReq) (models.{{Model}}ID, error) {
    id, err := u.db.Insert(&models.{{Model}}{
        {{Fields}}
    })

    return id, err
}`

const CONTROLLER_SHOW = `func (u *{{Controller}}) Show(c context.Context, id models.{{Model}}ID) (*models.{{Model}}, error) {
    return u.db.FindByID(id)
}`

const CONTROLLER_UPDATE = `func (u *{{Controller}}) Update(c context.Context, r *UpdateReq, id models.{{Model}}ID) error {
    return u.db.Update(&models.{{Model}}{
        ID:       id,
        {{Fields}}
    })
}`

const CONTROLLER_DELETE = `func (u *{{Controller}}) Delete(c context.Context, id models.{{Model}}ID) error {
    return u.db.Delete(id)
}`
