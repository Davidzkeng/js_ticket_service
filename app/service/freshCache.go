package service

import "js_ticket_service/app/utils/log"

func FreshCache()  {
	defer func() {
		if err := recover();err != nil{
			logger.Logger.Error()
		}
	}()
}